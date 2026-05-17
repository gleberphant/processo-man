package casosdeuso

import "github.com/gleberphant/ProcessoMan/internal/entidades"

func (u *CasosDeUsoProcesso) ListarProcessos() ([]entidades.Processo, error) {
	lista, err := u.RepoProcessos.Listar()
	if err != nil {
		return nil, err
	}
	return lista, nil
}
