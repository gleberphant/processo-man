package CasosDeUsoUsuario

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
)

func (u *CasosDeUsoUsuario) AtualizarUsuario(usuario entidades.Usuario) error {

	err := u.RepoUsuarios.Atualizar(usuario)

	return err
}
