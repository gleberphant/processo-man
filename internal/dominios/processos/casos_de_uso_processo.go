package processos

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/dominios/tarefas"
	"github.com/google/uuid"
)

type IRepositorioProcesso interface {
	Criar(Processo) error
	Listar() ([]Processo, error)
	Atualizar(Processo) error
	Deletar(uuid.UUID) error
	BuscarPorUUID(uuid.UUID) (*Processo, error)
}

type IRepositorioTarefa interface {
	CriarTarefa(tarefas.Tarefa) error
	ListarTarefas(uuid.UUID) ([]tarefas.Tarefa, error)
	EditarTarefa(tarefas.Tarefa) error
	DeletarTarefa(uuid.UUID) error
	BuscarTarefaPorUUID(UUID uuid.UUID) (*tarefas.Tarefa, error)
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
		processo.UUID = uuid.New()

		return u.repoProcesso.Criar(processo)

	}

	return u.repoProcesso.Atualizar(processo)

}

func (u *CDUProcesso) ListarProcessos() ([]Processo, error) {

	return u.repoProcesso.Listar()
}

func (u *CDUProcesso) EditarProcesso(processo Processo) error {

	if processo.UUID == uuid.Nil {
		return errors.New("não é possível atualizar um processo sem UUID")
	}

	return u.repoProcesso.Atualizar(processo)
}

func (u *CDUProcesso) DeletarProcesso(strUUID string) error {

	UUID, err := uuid.Parse(strUUID)

	if err != nil {
		return err
	}

	return u.repoProcesso.Deletar(UUID)
}

func (u *CDUProcesso) BuscarProcessoPorUUID(strUUID string) (*Processo, error) {

	if strUUID == "" {
		return nil, errors.New("UUID nulo")
	}

	var err error
	var processo *Processo
	var tarefas []tarefas.Tarefa

	processo, err = u.repoProcesso.BuscarPorUUID(uuid.MustParse(strUUID))

	if err != nil {
		return nil, err
	}

	tarefas, err = u.repoTarefa.ListarTarefas(processo.UUID)

	processo.Tarefas = tarefas

	return processo, nil

}
