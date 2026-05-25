package autenticacao

import (
	"errors"
	"fmt"

	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IRepositorioToken interface {
	Fechar()
	Criar(*Token) (*Token, error)
	BuscarPorUUID(UUID uuid.UUID) (*Token, error)
	DeletarPorUsuarioUUID(uuid.UUID) error
	VerificarPermissaoPerfil(string, string) bool
	//ObterMapaPermissoes(rota string) map[string]bool
}

type IRepositorioUsuario interface {
	Fechar()
	BuscarPorEmail(string) (*usuarios.Usuario, error)
	BuscarPorUUID(uuid.UUID) (*usuarios.Usuario, error)
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

// verificar se token tem permissão para acessar uma rota
func (a *CDUAutenticacao) VerificarPermissao(tokenUUID uuid.UUID, rota string) error {

	// busca o token completo
	token, err := a.RepoTokens.BuscarPorUUID(tokenUUID)

	if err != nil {
		return fmt.Errorf("Token não encontrado: %w ", err)
	}

	// verificar permissao
	for _, perfil := range token.Perfis {
		if a.RepoTokens.VerificarPermissaoPerfil(perfil, rota) {
			return nil
		}
	}

	return fmt.Errorf("perfil [%v] não autorizado a rota [%s]", token.Perfis, rota)

}

func (a *CDUAutenticacao) AutenticarUsuario(email string, senha string) (string, error) {

	//valida usuario
	if email == "" {
		return "", errors.New("Usuario nulo")
	}
	// 1 Verifica se o usuario/email existe
	usuario, err := a.RepoUsuarios.BuscarPorEmail(email)
	if err != nil {
		return "", fmt.Errorf("usuario não encontrado: %w ", err)
	}

	// 2 verifica senha
	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha))

	if err != nil {
		//fallback para usuario teste
		if senha != usuario.Senha && usuario.UUID.String() == "00000000-0000-0000-0000-000000000000" {
			return "", fmt.Errorf("senha inválida : %w ", err)
		}
	}

	// Gerar token
	token, err := a.GerarToken(usuario)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar token: %w ", err)
	}

	//retornar o uuid do token gerado
	return token.UUID.String(), nil
}

// gerar token
func (a *CDUAutenticacao) GerarToken(usuario *usuarios.Usuario) (*Token, error) {

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
		&Token{
			UUID:        tokenUUID,
			UsuarioUUID: usuario.UUID,
			Validade:    "temporario",
			Perfis:      usuario.Perfis,
		})

	if err != nil {
		return nil, err
	}

	return token, nil

}
