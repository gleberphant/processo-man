package tarefas

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

type IRepositorioTarefa interface {
	Fechar()
	CriarTarefa(entidades.Tarefa) error
	ListarTarefas() ([]entidades.Tarefa, error)
	ListarTarefasPorProcesso(uuid.UUID) ([]entidades.Tarefa, error)
	ListarTarefasPorResponsavel(uuid.UUID) ([]entidades.Tarefa, error)
	EditarTarefa(entidades.Tarefa) error
	DeletarTarefa(uuid.UUID) error
	BuscarTarefaPorUUID(UUID uuid.UUID) (*entidades.Tarefa, error)
}

type IRepositorioProcesso interface {
	Fechar()
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

func (a *CDUTarefa) Fechar() error {
	a.repoProcesso.Fechar()
	a.repoTarefa.Fechar()
	return nil
}

func (t *CDUTarefa) CriarTarefa(tarefa entidades.Tarefa) error {

	tarefa.UUID, _ = uuid.NewV7()

	if tarefa.ProcessoUUID == uuid.Nil {
		return errors.New("Informe o processo a qual pertence a tarefa ")
	}

	return t.repoTarefa.CriarTarefa(tarefa)

}

func (t *CDUTarefa) ListarTarefasPorProcesso(processoUUID uuid.UUID) ([]entidades.Tarefa, error) {

	return t.repoTarefa.ListarTarefasPorProcesso(processoUUID)
}

func (t *CDUTarefa) ListarTarefasPorResponsavel(responsavelUUID uuid.UUID) ([]entidades.Tarefa, error) {

	return t.repoTarefa.ListarTarefasPorResponsavel(responsavelUUID)
}

func (t *CDUTarefa) ListarTarefas() ([]entidades.Tarefa, error) {

	return t.repoTarefa.ListarTarefas()
}

func (t *CDUTarefa) EditarTarefa(tarefa entidades.Tarefa) error {

	return t.repoTarefa.EditarTarefa(tarefa)
}

func (t *CDUTarefa) DeletarTarefa(tarefaUUID uuid.UUID) error {

	return t.repoTarefa.DeletarTarefa(tarefaUUID)
}

func (t *CDUTarefa) BuscarTarefaPorUUID(tarefaUUID uuid.UUID) (*entidades.Tarefa, error) {
	return t.repoTarefa.BuscarTarefaPorUUID(tarefaUUID)
}

func (t *CDUTarefa) ValidarProcesso(processoUUID uuid.UUID) error {
	return t.repoProcesso.ValidarProcesso(processoUUID)
}
