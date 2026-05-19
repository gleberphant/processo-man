package entidades

import (
	"time"

	"github.com/google/uuid"
)

type Processo struct {
	UUID        uuid.UUID `json:"uuid,omitempty"  db:"uuid"`
	Nome        string    `json:"nome,omitempty"  db:"nome"`
	Dono        Cliente   `json:"dono,omitempty"  db:"dono"`
	DataCriacao time.Time `json:"data_criacao,omitempty"  db:"data_criacao"`
	Tarefas     []Tarefa  `json:"tarefas,omitempty"  db:""`
	Comentarios string    `json:"comentarios,omitempty"  db:"comentarios"`
}
