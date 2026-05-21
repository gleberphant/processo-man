package tarefas

import (
	"errors"

	"github.com/google/uuid"
)

type IRepositorioTarefa interface {
	CriarTarefa(Tarefa) error
	ListarTarefas(uuid.UUID) ([]Tarefa, error)
	EditarTarefa(Tarefa) error
	DeletarTarefa(uuid.UUID) error
	BuscarTarefaPorUUID(UUID uuid.UUID) (*Tarefa, error)
}

type CDUTarefa struct {
	repoTarefa IRepositorioTarefa
}

func NovoCDUTarefa(TarefasRepo IRepositorioTarefa) *CDUTarefa {

	return &CDUTarefa{
		repoTarefa: TarefasRepo,
	}
}

func (t *CDUTarefa) CriarTarefa(tarefa Tarefa) error {

	tarefa.UUID = uuid.New()

	if tarefa.ProcessoUUID == uuid.Nil {
		return errors.New("Informe o processo a qual pertence a tarefa ")
	}

	return t.repoTarefa.CriarTarefa(tarefa)

}

func (t *CDUTarefa) ListarTarefas(strProcessoUUID string) ([]Tarefa, error) {

	processoUUID, err := uuid.Parse(strProcessoUUID)

	if err != nil {
		return nil, err
	}

	return t.repoTarefa.ListarTarefas(processoUUID)
}

func (t *CDUTarefa) EditarTarefa(tarefa Tarefa) error {

	return nil
}

func (t *CDUTarefa) DeletarTarefa(strUUID string) error {

	return nil
}

func (t *CDUTarefa) BuscarTarefaPorUUID(strUUID string) (Tarefa, error) {

	return Tarefa{}, nil
}
