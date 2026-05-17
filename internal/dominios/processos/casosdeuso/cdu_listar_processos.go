package casosdeuso

import "github.com/gleberphant/ProcessoMan/internal/entidades"

func (u *CDUProcesso) ListarProcessos() ([]entidades.Processo, error) {

	return u.repoProcesso.Listar()
}
