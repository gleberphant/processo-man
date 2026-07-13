package servicos

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/aplicacao/repositorios"
	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"

	"github.com/google/uuid"
)

type ServicoTarefa struct {
	repoTarefa   *repositorios.RepositorioTarefa
	repoProcesso *repositorios.RepositorioProcesso
}

func NovoCDUTarefa(TarefasRepo *repositorios.RepositorioTarefa, ProcessoRepo *repositorios.RepositorioProcesso) *ServicoTarefa {

	return &ServicoTarefa{
		repoTarefa:   TarefasRepo,
		repoProcesso: ProcessoRepo,
	}
}

func (a *ServicoTarefa) Fechar() error {
	a.repoProcesso.Fechar()
	a.repoTarefa.Fechar()
	return nil
}

func (t *ServicoTarefa) CriarTarefa(tarefa entidades.Tarefa) error {

	tarefa.UUID, _ = uuid.NewV7()

	if tarefa.ProcessoUUID == uuid.Nil {
		return errors.New("Informe o processo a qual pertence a tarefa ")
	}

	return t.repoTarefa.CriarTarefa(tarefa)

}

func (t *ServicoTarefa) ListarTarefasPorProcesso(processoUUID uuid.UUID) ([]entidades.Tarefa, error) {

	return t.repoTarefa.ListarTarefasPorProcesso(processoUUID)
}

func (t *ServicoTarefa) ListarTarefasPorResponsavel(responsavelUUID uuid.UUID) ([]entidades.Tarefa, error) {

	return t.repoTarefa.ListarTarefasPorResponsavel(responsavelUUID)
}

func (t *ServicoTarefa) ListarTarefas() ([]entidades.Tarefa, error) {

	return t.repoTarefa.ListarTarefas()
}

func (t *ServicoTarefa) EditarTarefa(tarefa entidades.Tarefa) error {

	return t.repoTarefa.EditarTarefa(tarefa)
}

func (t *ServicoTarefa) DeletarTarefa(tarefaUUID uuid.UUID) error {

	return t.repoTarefa.DeletarTarefa(tarefaUUID)
}

func (t *ServicoTarefa) BuscarTarefaPorUUID(tarefaUUID uuid.UUID) (*entidades.Tarefa, error) {
	return t.repoTarefa.BuscarTarefaPorUUID(tarefaUUID)
}

func (t *ServicoTarefa) ValidarProcesso(processoUUID uuid.UUID) error {
	return t.repoProcesso.ValidarProcesso(processoUUID)
}
