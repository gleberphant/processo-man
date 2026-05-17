package casosdeuso

import (
	"github.com/google/uuid"
)

func (u *CDUProcesso) DeletarProcesso(strUUID string) error {

	UUID, err := uuid.Parse(strUUID)

	if err != nil {
		return err
	}

	return u.repoProcesso.Deletar(UUID)
}
