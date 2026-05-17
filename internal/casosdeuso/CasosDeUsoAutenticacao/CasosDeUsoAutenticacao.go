package CasosDeUsoAutenticacao

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
)

type IRepositorioToken interface {
	Fechar()
	Criar(entidades.Token) (*entidades.Token, error)
	BuscarPorUUID(entidades.Token) (*entidades.Token, error)
	DeletarPorUsuarioUUID(UsuarioUUID string) error
}

type IRepositorioUsuario interface {
	Fechar()
	AutenticarUsuario(string, string) (string, error)
}

type CasosDeUsoAutenticacao struct {
	RepoTokens   IRepositorioToken
	RepoUsuarios IRepositorioUsuario
}

func NovoCasoDeUsoAutenticacao(tokensRepo IRepositorioToken, usuariosRepo IRepositorioUsuario) *CasosDeUsoAutenticacao {

	return &CasosDeUsoAutenticacao{
		RepoTokens:   tokensRepo,
		RepoUsuarios: usuariosRepo,
	}
}
