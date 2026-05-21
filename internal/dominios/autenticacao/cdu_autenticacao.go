package autenticacao

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

type IRepositorioToken interface {
	Fechar()
	Criar(entidades.Token) (*entidades.Token, error)
	BuscarPorUUID(entidades.Token) (*entidades.Token, error)
	DeletarPorUsuarioUUID(UsuarioUUID string) error
	AutenticarUsuario(string, string) (string, error)
}

// type IRepositorioUsuario interface {
// 	Fechar()
// 	AutenticarUsuario(string, string) (string, error)
// }

type CDUAutenticacao struct {
	RepoTokens IRepositorioToken
	//RepoUsuarios IRepositorioUsuario
}

func NovoCDUAutenticacao(tokensRepo IRepositorioToken) *CDUAutenticacao {

	return &CDUAutenticacao{
		RepoTokens: tokensRepo,
		//RepoUsuarios: usuariosRepo,
	}
}

func (a *CDUAutenticacao) AutenticarUsuario(email string, senha string) (*entidades.Token, error) {

	// Verifica se o usuario existe

	if email == "" {
		return nil, errors.New("Usuario nulo")
	}

	usuarioUUID, err := a.RepoTokens.AutenticarUsuario(email, senha)

	if err != nil {
		return nil, err
	}

	// Gera token
	token, err := a.GerarToken(usuarioUUID)
	if err != nil {
		return nil, err
	}

	//retornar o token gerado
	return token, nil
}

// gerar token
func (a *CDUAutenticacao) GerarToken(usuarioUUID string) (*entidades.Token, error) {

	if usuarioUUID == "" {
		return nil, errors.New("usuário inválido")
	}

	// limpar tokens antigos no repositorio
	err := a.RepoTokens.DeletarPorUsuarioUUID(usuarioUUID)

	if err != nil {
		return nil, err
	}

	// inserir o novo token no repositorio
	token, err := a.RepoTokens.Criar(entidades.Token{
		UUID:        uuid.New(),
		UsuarioUUID: uuid.MustParse(usuarioUUID),
		Validade:    "temporario",
	})

	if err != nil {
		return nil, err
	}

	return token, nil

}

// verificar se token é valido. retorna error se token não encontrado
func (a *CDUAutenticacao) ValidarToken(token entidades.Token) error {

	if token.UUID == uuid.Nil {
		return errors.New("token inválido: UUID não fornecido")
	}

	_, err := a.RepoTokens.BuscarPorUUID(token)

	if err != nil {
		return err
	}

	return nil
}
