package casosdeuso

import "github.com/gleberphant/ProcessoMan/internal/entidades"

func (u *Usuario) ListarUsuarios() ([]entidades.Usuario, error) {
	lista, err := u.RepoUsuarios.Listar()
	if err != nil {
		return nil, err
	}
	return lista, nil
}
