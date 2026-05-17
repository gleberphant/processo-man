package casosdeuso

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
)

func (u *CasosDeUsoProcesso) AtualizarProcesso(processo entidades.Processo) error {

	err := u.RepoProcessos.Atualizar(processo)

	return err
}
