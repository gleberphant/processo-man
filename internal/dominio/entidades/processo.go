package entidades

import (
	"time"

	"github.com/google/uuid"
)

type Processo struct {
	UUID        uuid.UUID
	Nome        string
	Dono        Cliente
	Responsavel Colaborador
	DataCriacao time.Time
	Tarefas     []Tarefa
	Comentarios string
}
