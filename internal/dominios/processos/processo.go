package processos

import (
	"time"

	"github.com/gleberphant/ProcessoMan/internal/dominios/tarefas"
	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"

	"github.com/google/uuid"
)

type Processo struct {
	UUID        uuid.UUID
	Nome        string
	Dono        usuarios.Cliente
	DataCriacao time.Time
	Tarefas     []tarefas.Tarefa
	Comentarios string
}
