package entidades

import (
	"time"

	"github.com/google/uuid"
)

type Tarefa struct {
	UUID          uuid.UUID   `json:"uuid,omitempty"  db:"uuid"`
	Responsavel   Colaborador `json:"responsavel,omitempty"  db:"responsavel"`
	Nome          string      `json:"descricao,omitempty"  db:"descricao"`
	Conclusao     bool        `json:"conclusao,omitempty"  db:"conclusao"`
	DataConclusao time.Time   `json:"data_conclusao,omitempty"  db:"data_conclusao"`
	DataCriacao   time.Time   `json:"data_criacao,omitempty"  db:"data_criacao"`
}
