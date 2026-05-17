package casosdeuso

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

func (u *CDUUsuario) BuscarUsuarioPorUUID(strUUID string) (*entidades.Usuario, error) {

	if strUUID == "" {
		return nil, errors.New("UUID nulo")
	}

	return u.RepoUsuarios.BuscarPorUUID(uuid.MustParse(strUUID))
}
