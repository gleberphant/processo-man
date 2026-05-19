package entidades

import (
	"time"

	"github.com/google/uuid"
)

type Tarefa struct {
	UUID            uuid.UUID `json:"uuid,omitempty"  db:"uuid"`
	ProcessoUUID    uuid.UUID `json:"processo_id,omitempty"  db:"processo_id"`
	ResponsavelUUID uuid.UUID `json:"responsavel,omitempty"  db:"responsavel"`
	Nome            string    `json:"descricao,omitempty"  db:"descricao"`
	Conclusao       bool      `json:"conclusao,omitempty"  db:"conclusao"`
	Comentarios     string    `json:"comentarios,omitempty"  db:"conclusao"`
	DataConclusao   time.Time `json:"data_conclusao,omitempty"  db:"data_conclusao"`
	DataCriacao     time.Time `json:"data_criacao,omitempty"  db:"data_criacao"`
}
