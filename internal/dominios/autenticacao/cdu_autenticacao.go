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
	Criar(Token) (*Token, error)
	BuscarPorUUID(UUID uuid.UUID) (*Token, error)
	DeletarPorUsuarioUUID(uuid.UUID) error
}

type IRepositorioUsuario interface {
	Fechar()
	BuscarPorEmail(email string) (*usuarios.Usuario, error)
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

// verificar se token é valido. retorna error se token não encontrado
func (a *CDUAutenticacao) ValidarToken(strUUID string) error {

	UUID, err := uuid.Parse(strUUID)
	if err != nil {
		return fmt.Errorf("token nulo: %w ", err)

	}

	_, err = a.RepoTokens.BuscarPorUUID(UUID)

	if err != nil {
		return fmt.Errorf("token invalido: %w ", err)
	}

	return nil
}

func (a *CDUAutenticacao) AutenticarUsuario(email string, senha string) (string, error) {

	//valida usuario
	// 1 Verifica se o usuario/email existe
	if email == "" {
		return "", errors.New("Usuario nulo")
	}

	usuario, err := a.RepoUsuarios.BuscarPorEmail(email)
	if err != nil {
		return "", fmt.Errorf("usuario não encontrado: %w ", err)
	}

	// 2 verifica senha
	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha))

	if err != nil {
		if senha != usuario.Senha {
			return "", fmt.Errorf("senha inválida : %w ", err)
		}
		//return "", fmt.Errorf("senha inválida : %w ", err)
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

	// gerar novo token
	tokenUUID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	// inserir o novo token no repositorio
	token, err := a.RepoTokens.Criar(Token{
		UUID:        tokenUUID,
		UsuarioUUID: usuario.UUID,
		Validade:    "temporario",
	})

	if err != nil {
		return nil, err
	}

	return token, nil

}
