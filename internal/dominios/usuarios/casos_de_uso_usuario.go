package usuarios

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

type IRepositorioUsuario interface {
	Criar(Usuario) error
	Listar() ([]Usuario, error)
	Atualizar(Usuario) error
	Deletar(uuid.UUID) error
	BuscarPorUUID(uuid.UUID) (*Usuario, error)
}

type CDUUsuario struct {
	RepoUsuarios IRepositorioUsuario
}

func NovoCDUUsuario(usuariosRepo IRepositorioUsuario) *CDUUsuario {

	return &CDUUsuario{
		RepoUsuarios: usuariosRepo,
	}
}

func (u *CDUUsuario) CriaUsuario(usuario Usuario) error {

	usuario.UUID = uuid.New()

	for i := range usuario.Perfis {
		usuario.Perfis[i].UUID = uuid.New()
	}

	log.Printf("Criando Usuario %v", usuario)
	err := u.RepoUsuarios.Criar(usuario)

	return err

}

func (u *CDUUsuario) ListarUsuarios() ([]Usuario, error) {
	lista, err := u.RepoUsuarios.Listar()
	if err != nil {
		return nil, err
	}
	return lista, nil
}

func (u *CDUUsuario) EditarUsuario(usuario Usuario) error {

	err := u.RepoUsuarios.Atualizar(usuario)

	return err
}

func (u *CDUUsuario) DeletarUsuario(strUUID string) error {

	UUID, err := uuid.Parse(strUUID)
	if err != nil {
		return err
	}

	u.RepoUsuarios.Deletar(UUID)

	return nil
}

func (u *CDUUsuario) BuscarUsuarioPorUUID(strUUID string) (*Usuario, error) {

	if strUUID == "" {
		return nil, errors.New("UUID nulo")
	}

	return u.RepoUsuarios.BuscarPorUUID(uuid.MustParse(strUUID))
}
