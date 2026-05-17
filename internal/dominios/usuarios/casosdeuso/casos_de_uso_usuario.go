package casosdeuso

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

type IRepositorioUsuario interface {
	Criar(entidades.Usuario) error
	Listar() ([]entidades.Usuario, error)
	Atualizar(entidades.Usuario) error
	Deletar(uuid.UUID) error
	BuscarPorUUID(uuid.UUID) (*entidades.Usuario, error)
}

type CDUUsuario struct {
	RepoUsuarios IRepositorioUsuario
}

func NovoCDUUsuario(usuariosRepo IRepositorioUsuario) *CDUUsuario {

	return &CDUUsuario{
		RepoUsuarios: usuariosRepo,
	}
}
