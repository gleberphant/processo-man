package CasosDeUsoUsuario

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

type CasosDeUsoUsuario struct {
	RepoUsuarios IRepositorioUsuario
}

func NovoCasoDeUsoUsuario(usuariosRepo IRepositorioUsuario) *CasosDeUsoUsuario {

	return &CasosDeUsoUsuario{
		RepoUsuarios: usuariosRepo,
	}
}
