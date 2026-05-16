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

type Tarefa struct {
	UUID          uuid.UUID   `json:"uuid,omitempty"  db:"uuid"`
	Responsavel   Colaborador `json:"responsavel,omitempty"  db:"responsavel"`
	Descricao     string      `json:"descricao,omitempty"  db:"descricao"`
	Conclusao     bool        `json:"conclusao,omitempty"  db:"conclusao"`
	DataConclusao time.Time   `json:"data_conclusao,omitempty"  db:"data_conclusao"`
	DataCriacao   time.Time   `json:"data_criacao,omitempty"  db:"data_criacao"`
}

type Cliente struct {
	Usuario `json:"usuario,omitempty"  db:"usuario"`
}

type Colaborador struct {
	Usuario `json:"usuario,omitempty"  db:"usuario"`
}
