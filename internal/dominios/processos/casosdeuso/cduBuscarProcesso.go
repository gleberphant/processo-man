package casosdeuso

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

func (u *CasosDeUsoProcesso) BuscarUsuarioPorUUID(strUUID string) (*entidades.Processo, error) {

	if strUUID == "" {
		return nil, errors.New("UUID nulo")
	}
	usuario, err := u.RepoProcessos.BuscarPorUUID(uuid.MustParse(strUUID))
	if err != nil {
		return nil, err
	}

	return usuario, nil
}
