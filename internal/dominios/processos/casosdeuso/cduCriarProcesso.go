package casosdeuso

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

func (u *CasosDeUsoProcesso) CriaProcesso(processo entidades.Processo) error {

	if processo.UUID == uuid.Nil {
		processo.UUID = uuid.New()

		for i := range processo.ListaTarefas {
			processo.ListaTarefas[i].UUID = uuid.New()
		}
		err := u.RepoProcessos.Criar(processo)

		return err

	} else {

		err := u.RepoProcessos.Atualizar(processo)

		return err

	}

}
