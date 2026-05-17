package CasosDeUsoUsuario

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

func (u *CasosDeUsoUsuario) BuscarUsuarioPorUUID(strUUID string) (*entidades.Usuario, error) {

	if strUUID == "" {
		return nil, errors.New("UUID nulo")
	}
	usuario, err := u.RepoUsuarios.BuscarPorUUID(uuid.MustParse(strUUID))
	if err != nil {
		return nil, err
	}

	return usuario, nil
}
