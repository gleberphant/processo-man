package tarefas

import (
	"errors"

	"github.com/google/uuid"
)

type IRepositorioTarefa interface {
	CriarTarefa(Tarefa) error
	ListarTarefas() ([]Tarefa, error)
	ListarTarefasPorProcesso(uuid.UUID) ([]Tarefa, error)
	ListarTarefasPorResponsavel(uuid.UUID) ([]Tarefa, error)
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

func (t *CDUTarefa) ListarTarefasPorProcesso(processoUUID uuid.UUID) ([]Tarefa, error) {

	return t.repoTarefa.ListarTarefasPorProcesso(processoUUID)
}

func (t *CDUTarefa) ListarTarefasPorResponsavel(responsavelUUID uuid.UUID) ([]Tarefa, error) {

	return t.repoTarefa.ListarTarefasPorResponsavel(responsavelUUID)
}

func (t *CDUTarefa) ListarTarefas() ([]Tarefa, error) {

	return t.repoTarefa.ListarTarefas()
}

func (t *CDUTarefa) EditarTarefa(tarefa Tarefa) error {

	return t.repoTarefa.EditarTarefa(tarefa)
}

func (t *CDUTarefa) DeletarTarefa(tarefaUUID uuid.UUID) error {

	return t.repoTarefa.DeletarTarefa(tarefaUUID)
}

func (t *CDUTarefa) BuscarTarefaPorUUID(tarefaUUID uuid.UUID) (*Tarefa, error) {
	return t.repoTarefa.BuscarTarefaPorUUID(tarefaUUID)
}

func (t *CDUTarefa) ValidarProcesso(processoUUID uuid.UUID) error {
	return t.repoProcesso.ValidarProcesso(processoUUID)
}
