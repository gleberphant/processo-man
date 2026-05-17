package casosdeuso

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

func (u *CDUUsuario) CriaUsuario(usuario entidades.Usuario) error {

	if usuario.UUID == uuid.Nil {
		usuario.UUID = uuid.New()

		for i := range usuario.Perfis {
			usuario.Perfis[i].UUID = uuid.New()
		}
		err := u.RepoUsuarios.Criar(usuario)

		return err

	} else {

		err := u.RepoUsuarios.Atualizar(usuario)

		return err

	}

}
