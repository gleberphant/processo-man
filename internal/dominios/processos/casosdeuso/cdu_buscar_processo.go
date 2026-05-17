package casosdeuso

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

func (u *CDUProcesso) BuscarProcessoPorUUID(strUUID string) (*entidades.Processo, error) {

	if strUUID == "" {
		return nil, errors.New("UUID nulo")
	}

	return u.repoProcesso.BuscarPorUUID(uuid.MustParse(strUUID))

}
