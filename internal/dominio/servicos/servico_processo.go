package servicos

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/aplicacao/repositorios"
	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/google/uuid"
)

type CDUProcesso struct {
	repoProcesso *repositorios.RepositorioProcesso
	repoTarefa   *repositorios.RepositorioTarefa
}

func NovoCDUProcesso(ProcessosRepo *repositorios.RepositorioProcesso, TarefasRepo *repositorios.RepositorioTarefa) *CDUProcesso {

	return &CDUProcesso{
		repoProcesso: ProcessosRepo,
		repoTarefa:   TarefasRepo,
	}
}

func (a *CDUProcesso) Fechar() error {
	a.repoProcesso.Fechar()
	a.repoTarefa.Fechar()
	return nil
}

func (u *CDUProcesso) CriarProcesso(processo entidades.Processo) error {

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

func (u *CDUProcesso) ListarProcessos() ([]entidades.Processo, error) {

	return u.repoProcesso.Listar()

}

func (u *CDUProcesso) ListarProcessosPorCliente(UUID uuid.UUID) ([]entidades.Processo, error) {

	return u.repoProcesso.ListarProcessosPorCliente(UUID)

}

func (u *CDUProcesso) EditarProcesso(processo entidades.Processo) error {

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

func (u *CDUProcesso) BuscarProcessoPorUUID(processoUUID uuid.UUID) (*entidades.Processo, error) {

	if processoUUID == uuid.Nil {
		return nil, errors.New("UUID nulo")
	}

	var err error
	var processo *entidades.Processo
	var tarefas []entidades.Tarefa

	processo, err = u.repoProcesso.BuscarPorUUID(processoUUID)

	if err != nil {
		return nil, err
	}

	tarefas, err = u.repoTarefa.ListarTarefasPorProcesso(processo.UUID)

	processo.Tarefas = tarefas

	return processo, nil

}
