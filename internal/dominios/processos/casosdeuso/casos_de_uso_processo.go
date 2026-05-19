package casosdeuso

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

type IRepositorioProcesso interface {
	Criar(entidades.Processo) error
	Listar() ([]entidades.Processo, error)
	Atualizar(entidades.Processo) error
	Deletar(uuid.UUID) error
	BuscarPorUUID(uuid.UUID) (*entidades.Processo, error)
}

type IRepositorioTarefa interface {
	CriarTarefa(entidades.Tarefa) error
	ListarTarefas(uuid.UUID) ([]entidades.Tarefa, error)
	AtualizarTarefa(entidades.Tarefa) error
	DeletarTarefa(uuid.UUID) error
	BuscarTarefaPorUUID(UUID uuid.UUID) (*entidades.Tarefa, error)
}

type CDUProcesso struct {
	repoProcesso IRepositorioProcesso
	repoTarefa   IRepositorioTarefa
}

func NovoCDUProcesso(ProcessosRepo IRepositorioProcesso, TarefasRepo IRepositorioTarefa) *CDUProcesso {

	return &CDUProcesso{
		repoProcesso: ProcessosRepo,
		repoTarefa:   TarefasRepo,
	}
}
