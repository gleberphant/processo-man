package casosdeuso

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

func (u *CDUProcesso) CriarProcesso(processo entidades.Processo) error {

	for i := range processo.ListaTarefas {
		if processo.ListaTarefas[i].UUID == uuid.Nil {
			processo.ListaTarefas[i].UUID = uuid.New()
		}
	}

	if processo.UUID == uuid.Nil {
		processo.UUID = uuid.New()

		return u.repoProcesso.Criar(processo)

	}

	return u.repoProcesso.Atualizar(processo)

}
