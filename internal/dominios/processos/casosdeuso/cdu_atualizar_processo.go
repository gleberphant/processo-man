package casosdeuso

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

func (u *CDUProcesso) AtualizarProcesso(processo entidades.Processo) error {

	if processo.UUID == uuid.Nil {
		return errors.New("não é possível atualizar um processo sem UUID")
	}

	return u.repoProcesso.Atualizar(processo)
}
