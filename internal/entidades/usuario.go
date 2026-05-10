package entidades

import (
	"time"

	"github.com/google/uuid" // Certifique-se de ter rodado 'go get' para este pacote
)

type Usuario struct {
	ID          uuid.UUID `json:"id,omitempty"  db:"id"`
	Nome        string    `json:"nome,omitempty"  db:"nome"`
	Email       string    `json:"email,omitempty"  db:"email"`
	Senha       string    `json:"-"  db:"senha"`
	DataCriacao time.Time `json:"data_criacao,omitempty"  db:"data_criacao"`
	Perfis      []Perfil  `json:"perfis,omitempty"  db:"-"`
}
