package casosdeuso

import (
	"github.com/google/uuid"
)

func (u *CasosDeUsoProcesso) DeletarUsuario(strUUID string) error {

	UUID, err := uuid.Parse(strUUID)
	if err != nil {
		return err
	}

	u.RepoProcessos.Deletar(UUID)

	return nil
}
