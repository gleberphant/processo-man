package casosdeuso

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

func (u *CDUProcesso) CriarTarefa(tarefa entidades.Tarefa) error {

	if tarefa.UUID == uuid.Nil {
		tarefa.UUID = uuid.New()
	}

	if tarefa.ProcessoUUID == uuid.Nil {
		return errors.New("Informe o processo a qual pertence a tarefa ")
	}

	return u.repoTarefa.CriarTarefa(tarefa)

}
