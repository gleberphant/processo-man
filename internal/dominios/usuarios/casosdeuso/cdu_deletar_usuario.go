package casosdeuso

import (
	"github.com/google/uuid"
)

func (u *CDUUsuario) DeletarUsuario(strUUID string) error {

	UUID, err := uuid.Parse(strUUID)
	if err != nil {
		return err
	}

	u.RepoUsuarios.Deletar(UUID)

	return nil
}
