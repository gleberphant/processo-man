package processos

import (
	"time"

	"github.com/gleberphant/ProcessoMan/internal/dominios/tarefas"
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

type Processo struct {
	UUID        uuid.UUID
	Nome        string
	Dono        entidades.Cliente
	DataCriacao time.Time
	Tarefas     []tarefas.Tarefa
	Comentarios string
}
