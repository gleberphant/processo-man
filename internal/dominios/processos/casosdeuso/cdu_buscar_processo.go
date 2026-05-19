package casosdeuso

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

func (u *CDUProcesso) BuscarProcessoPorUUID(strUUID string) (*entidades.Processo, error) {

	if strUUID == "" {
		return nil, errors.New("UUID nulo")
	}

	var err error
	var processo *entidades.Processo
	var tarefas []entidades.Tarefa

	processo, err = u.repoProcesso.BuscarPorUUID(uuid.MustParse(strUUID))

	if err != nil {
		return nil, err
	}

	tarefas, err = u.repoTarefa.ListarTarefas(processo.UUID)

	processo.Tarefas = tarefas

	return processo, nil

}
