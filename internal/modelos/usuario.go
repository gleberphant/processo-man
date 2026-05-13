package modelos

import (
	"time"
)

type Usuario struct {
	UUID        string    `json:"uuid,omitempty"  db:"uuid"`
	Nome        string    `json:"nome,omitempty"  db:"nome"`
	Email       string    `json:"email,omitempty"  db:"email"`
	Senha       string    `json:"-"  db:"senha"`
	DataCriacao time.Time `json:"data_criacao,omitempty"  db:"data_criacao"`
	Perfis      []Perfil  `json:"perfis,omitempty"  db:"-"`
}
