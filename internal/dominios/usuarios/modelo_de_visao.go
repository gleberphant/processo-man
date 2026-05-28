package usuarios

import (
	"time"

	"github.com/google/uuid"
)

type ViewModelUsuario struct {
	UUID     string      `json:"uuid,omitempty"`
	Usuarios interface{} `json:"usuarios,omitempty"`
	Tarefas  interface{} `json:"tarefas,omitempty"`
	Anexos   interface{} `json:"anexos,omitempty"`
}

type tarefasView struct {
	UUID            uuid.UUID
	ProcessoUUID    uuid.UUID
	ResponsavelUUID uuid.UUID
	Nome            string
	Concluida       bool
	Comentarios     string
	DataConclusao   time.Time
	DataCriacao     time.Time
}
