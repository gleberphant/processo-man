package autenticacao

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IRepositorioToken interface {
	Fechar()
	Criar(*entidades.Token) (*entidades.Token, error)
	BuscarPorUUID(UUID uuid.UUID) (*entidades.Token, error)
	DeletarPorUsuarioUUID(uuid.UUID) error
	VerificarPermissaoPerfil(string, string) bool
	//ObterMapaPermissoes(rota string) map[string]bool
}

type IRepositorioUsuario interface {
	Fechar()
	BuscarPorEmail(string) (*entidades.Usuario, error)
	BuscarPorUUID(uuid.UUID) (*entidades.Usuario, error)
}

type CDUAutenticacao struct {
	RepoTokens   IRepositorioToken
	RepoUsuarios IRepositorioUsuario
}

func NovoCDUAutenticacao(tokensRepo IRepositorioToken, usuariosRepo IRepositorioUsuario) *CDUAutenticacao {

	return &CDUAutenticacao{
		RepoTokens:   tokensRepo,
		RepoUsuarios: usuariosRepo,
	}
}

func (a *CDUAutenticacao) Fechar() error {
	a.RepoTokens.Fechar()
	a.RepoUsuarios.Fechar()
	return nil
}

// verificar se token tem permissão para acessar uma rota
func (a *CDUAutenticacao) VerificarExisteToken(tokenUUID uuid.UUID) (*entidades.Token, error) {

	// Verifica se o token Existe
	token, err := a.RepoTokens.BuscarPorUUID(tokenUUID)
	if err != nil {
		return nil, fmt.Errorf("Token não encontrado: %w ", err)
	}

	return token, nil

}

// verificar se token tem permissão para acessar uma rota
func (a *CDUAutenticacao) VerificarPermissao(token *entidades.Token, rota string, metodo string) error {

	// verificar a permissão do permissao
	chaveRota := strings.ToLower(metodo + ":" + rota)

	for _, perfil := range token.Perfis {
		if a.RepoTokens.VerificarPermissaoPerfil(chaveRota, strings.ToLower(perfil)) {
			return nil
		}
	}

	return fmt.Errorf("perfil [%v] não autorizado a rota [%s]", token.Perfis, chaveRota)

}

/* função para autenticar o usuario, recebe o login e a senha, verifica se o usuario existe criar um um token associado a ele */
func (a *CDUAutenticacao) AutenticarUsuario(email string, senha string) (string, error) {

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
		if senha != usuario.Senha && usuario.UUID.String() == "00000000-0000-0000-0000-000000000000" {
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
func (a *CDUAutenticacao) gerarToken(usuario *entidades.Usuario) (*entidades.Token, error) {

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
