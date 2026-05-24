package tarefas

import (
	"errors"

	"github.com/google/uuid"
)

type IRepositorioTarefa interface {
	CriarTarefa(Tarefa) error
	ListarTarefas() ([]Tarefa, error)
	ListarTarefasPorProcesso(uuid.UUID) ([]Tarefa, error)
	EditarTarefa(Tarefa) error
	DeletarTarefa(uuid.UUID) error
	BuscarTarefaPorUUID(UUID uuid.UUID) (*Tarefa, error)
}

type IRepositorioProcesso interface {
	ValidarProcesso(uuid.UUID) error
}

type CDUTarefa struct {
	repoTarefa   IRepositorioTarefa
	repoProcesso IRepositorioProcesso
}

func NovoCDUTarefa(TarefasRepo IRepositorioTarefa, ProcessoRepo IRepositorioProcesso) *CDUTarefa {

	return &CDUTarefa{
		repoTarefa:   TarefasRepo,
		repoProcesso: ProcessoRepo,
	}
}

func (t *CDUTarefa) CriarTarefa(tarefa Tarefa) error {

	tarefa.UUID = uuid.New()

	if tarefa.ProcessoUUID == uuid.Nil {
		return errors.New("Informe o processo a qual pertence a tarefa ")
	}

	return t.repoTarefa.CriarTarefa(tarefa)

}

func (t *CDUTarefa) ListarTarefasPorProcesso(strProcessoUUID string) ([]Tarefa, error) {

	processoUUID, err := uuid.Parse(strProcessoUUID)

	if err != nil {
		return nil, err
	}

	return t.repoTarefa.ListarTarefasPorProcesso(processoUUID)
}

func (t *CDUTarefa) ListarTarefas() ([]Tarefa, error) {

	return t.repoTarefa.ListarTarefas()
}

func (t *CDUTarefa) EditarTarefa(tarefa Tarefa) error {

	return t.repoTarefa.EditarTarefa(tarefa)
}

func (t *CDUTarefa) DeletarTarefa(strUUID string) error {

	UUID, err := uuid.Parse(strUUID)

	if err != nil {
		return err
	}

	return t.repoTarefa.DeletarTarefa(UUID)
}

func (t *CDUTarefa) BuscarTarefaPorUUID(strUUID string) (*Tarefa, error) {
	UUID, err := uuid.Parse(strUUID)

	if err != nil {

		return nil, err
	}

	return t.repoTarefa.BuscarTarefaPorUUID(UUID)
}

func (t *CDUTarefa) ValidarProcesso(strUUID string) error {
	UUID, err := uuid.Parse(strUUID)

	if err != nil {

		return err
	}

	return t.repoProcesso.ValidarProcesso(UUID)
}
