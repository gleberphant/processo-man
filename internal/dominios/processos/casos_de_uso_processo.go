package processos

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/dominios/tarefas"
	"github.com/google/uuid"
)

type IRepositorioProcesso interface {
	Criar(Processo) error
	Listar() ([]Processo, error)
	Editar(Processo) error
	Deletar(uuid.UUID) error
	BuscarPorUUID(uuid.UUID) (*Processo, error)
}

type IRepositorioTarefa interface {
	CriarTarefa(tarefas.Tarefa) error
	ListarTarefas() ([]tarefas.Tarefa, error)
	ListarTarefasPorProcesso(uuid.UUID) ([]tarefas.Tarefa, error)
	EditarTarefa(tarefas.Tarefa) error
	DeletarTarefa(uuid.UUID) error
	BuscarTarefaPorUUID(UUID uuid.UUID) (*tarefas.Tarefa, error)
	DeletarTarefasPorProcesso(UUID uuid.UUID) error
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

func (u *CDUProcesso) CriarProcesso(processo Processo) error {

	for i := range processo.Tarefas {
		if processo.Tarefas[i].UUID == uuid.Nil {
			processo.Tarefas[i].UUID = uuid.New()
		}
	}

	if processo.UUID == uuid.Nil {
		processo.UUID, _ = uuid.NewV7()

		return u.repoProcesso.Criar(processo)

	}

	return u.repoProcesso.Editar(processo)

}

func (u *CDUProcesso) ListarProcessos() ([]Processo, error) {

	return u.repoProcesso.Listar()
}

func (u *CDUProcesso) EditarProcesso(processo Processo) error {

	if processo.UUID == uuid.Nil {
		return errors.New("não é possível atualizar um processo sem UUID")
	}

	return u.repoProcesso.Editar(processo)
}

func (u *CDUProcesso) DeletarProcesso(processoUUID uuid.UUID) error {

	err := u.repoProcesso.Deletar(processoUUID)
	if err != nil {
		return err
	}

	err = u.repoTarefa.DeletarTarefasPorProcesso(processoUUID)

	if err != nil {
		return err
	}

	return err
}

func (u *CDUProcesso) BuscarProcessoPorUUID(processoUUID uuid.UUID) (*Processo, error) {

	if processoUUID == uuid.Nil {
		return nil, errors.New("UUID nulo")
	}

	var err error
	var processo *Processo
	var tarefas []tarefas.Tarefa

	processo, err = u.repoProcesso.BuscarPorUUID(processoUUID)

	if err != nil {
		return nil, err
	}

	tarefas, err = u.repoTarefa.ListarTarefasPorProcesso(processo.UUID)

	processo.Tarefas = tarefas

	return processo, nil

}
