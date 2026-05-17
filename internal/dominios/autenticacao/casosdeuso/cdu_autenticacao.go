package casosdeuso

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
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
