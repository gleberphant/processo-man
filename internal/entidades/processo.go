package entidades

import (
	"time"

	"github.com/google/uuid"
)

type Processo struct {
	UUID         uuid.UUID `json:"uuid,omitempty"  db:"uuid"`
	Nome         string    `json:"nome,omitempty"  db:"nome"`
	Dono         Cliente   `json:"dono,omitempty"  db:"dono"`
	DataCriacao  time.Time `json:"data_criacao,omitempty"  db:"data_criacao"`
	ListaTarefas []Tarefa  `json:"lista_tarefas,omitempty"  db:"lista_tarefas"`
}
