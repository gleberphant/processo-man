package servicos

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gleberphant/ProcessoMan/internal/aplicacao/repositorios"
	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ServicoAutenticacao struct {
	RepoTokens   *repositorios.RepositorioTokenBolt
	RepoUsuarios *repositorios.RepositorioUsuario
}

func NovoCDUAutenticacao(tokensRepo *repositorios.RepositorioTokenBolt, usuariosRepo *repositorios.RepositorioUsuario) *ServicoAutenticacao {

	return &ServicoAutenticacao{
		RepoTokens:   tokensRepo,
		RepoUsuarios: usuariosRepo,
	}
}

func (a *ServicoAutenticacao) Fechar() error {
	a.RepoTokens.Fechar()
	a.RepoUsuarios.Fechar()
	return nil
}

// verificar se token tem permissão para acessar uma rota
func (a *ServicoAutenticacao) VerificarExisteToken(tokenUUID uuid.UUID) (*entidades.Token, error) {

	// Verifica se o token Existe
	token, err := a.RepoTokens.BuscarPorUUID(tokenUUID)
	if err != nil {
		return nil, fmt.Errorf("Token não encontrado: %w ", err)
	}

	return token, nil

}

// verificar se token tem permissão para acessar uma rota
func (a *ServicoAutenticacao) VerificarPermissao(token *entidades.Token, rota string, metodo string) error {
	var err error
	// verificar a permissão do permissao
	chaveRota := strings.ToLower(metodo + ":" + rota)

	for _, perfil := range token.Perfis {

		//verifica permissão
		err = a.RepoTokens.VerificarPermissaoPerfil(chaveRota, strings.ToLower(perfil))

		// se retornou sem erro é porque é autorizado
		if err == nil {
			return nil
		}

		// by passa para administrador - temporario
		if perfil == "Administrador" {
			log.Printf("** byPass Administrador %s : %s", token.UUID, token.UsuarioNome)
			return nil
		}

		log.Printf("Error: %v", err)
	}

	return fmt.Errorf(" Perfis [%v] não autorizados para a rota [%s]", token.Perfis, chaveRota)

}

/* função para autenticar o usuario, recebe o login e a senha, verifica se o usuario existe criar um um token associado a ele */
func (a *ServicoAutenticacao) AutenticarUsuario(email string, senha string) (string, error) {

	// valida formato do login
	if email == "" {
		return "", errors.New("Usuario nulo")
	}

	// 1. Verifica se o usuario/email existe
	usuario, err := a.RepoUsuarios.BuscarPorEmail(email)
	if err != nil {
		return "", fmt.Errorf("usuario não encontrado: %w ", err)
	}

	// 2. Verifica se a senha é correta
	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha))

	if err != nil {
		// DEV fallback para usuario teste
		if usuario.UUID.String() == "00000000-0000-0000-0000-000000000000" && senha == usuario.Senha {
			log.Printf("[ Logando com usuario teste %s perfis %s]\n ", usuario.Email, usuario.Perfis)
			//usuario de testes
		} else { // erro autenticação
			return "", fmt.Errorf("senha inválida : %w ", err)
		}
	}

	// Gerar token
	token, err := a.gerarToken(usuario)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar token: %w ", err)
	}

	//retornar o uuid do token gerado
	return token.UUID.String(), nil
}

// gerar token
func (a *ServicoAutenticacao) gerarToken(usuario *entidades.Usuario) (*entidades.Token, error) {

	// limpar tokens antigos no repositorio
	err := a.RepoTokens.DeletarPorUsuarioUUID(usuario.UUID)

	if err != nil {
		return nil, err
	}

	// gerar novo uuid do token
	tokenUUID, err := uuid.NewV7()

	if err != nil {
		return nil, err
	}

	// inserir o novo token no repositorio
	token, err := a.RepoTokens.Criar(
		&entidades.Token{
			UUID:        tokenUUID,
			UsuarioUUID: usuario.UUID,
			UsuarioNome: usuario.Nome,
			Validade:    "temporario",
			Perfis:      usuario.Perfis,
		})

	if err != nil {
		return nil, err
	}

	return token, nil

}
