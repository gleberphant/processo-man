package CasosDeUsoUsuario

import (
	"github.com/google/uuid"
)

func (u *CasosDeUsoUsuario) DeletarUsuario(strUUID string) error {

	UUID, err := uuid.Parse(strUUID)
	if err != nil {
		return err
	}

	u.RepoUsuarios.Deletar(UUID)

	return nil
}
