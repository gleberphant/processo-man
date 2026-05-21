package tarefas

import (
	"time"

	"github.com/google/uuid"
)

type Tarefa struct {
	UUID            uuid.UUID
	ProcessoUUID    uuid.UUID
	ResponsavelUUID uuid.UUID
	Nome            string
	Conclusao       bool
	Comentarios     string
	DataConclusao   time.Time
	DataCriacao     time.Time
}
