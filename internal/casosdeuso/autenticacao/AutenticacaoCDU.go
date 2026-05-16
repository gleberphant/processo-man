package autenticacao

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
)

type IRepositorioToken interface {
	Fechar()
	Deletar(entidades.Token) error
	Criar(entidades.Token) error
	BuscarPorUUID(entidades.Token) (*entidades.Token, error)
}

type IRepositorioUsuario interface {
	Fechar()
	AutenticarUsuario(*entidades.Usuario) error
}

type AutenticacaoCDU struct {
	RepoTokens   IRepositorioToken
	RepoUsuarios IRepositorioUsuario
}

func NovoAutenticacaoCDU(tokensRepo IRepositorioToken, usuariosRepo IRepositorioUsuario) *AutenticacaoCDU {

	return &AutenticacaoCDU{
		RepoTokens:   tokensRepo,
		RepoUsuarios: usuariosRepo,
	}
}
