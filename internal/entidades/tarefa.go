package entidades

import (
	"time"

	"github.com/google/uuid"
)

type Tarefa struct {
	UUID            uuid.UUID
	ProcessoUUID    uuid.UUID
	ResponsavelUUID uuid.UUID
	Nome            string
	Concluida       bool
	Comentarios     string
	DataConclusao   time.Time
	DataCriacao     time.Time
}
