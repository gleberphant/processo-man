package casosdeuso

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
)

func (u *CDUUsuario) EditarUsuario(usuario entidades.Usuario) error {

	err := u.RepoUsuarios.Atualizar(usuario)

	return err
}
