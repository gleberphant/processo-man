package modelos

import (
	"time"

	"github.com/google/uuid"
)

type Usuario struct {
	UUID        string    `json:"uuid,omitempty"  db:"uuid"`
	Nome        string    `json:"nome,omitempty"  db:"nome"`
	Email       string    `json:"email,omitempty"  db:"email"`
	Senha       string    `json:"-"  db:"senha"`
	DataCriacao time.Time `json:"data_criacao,omitempty"  db:"data_criacao"`
	Perfis      []Perfil  `json:"perfis,omitempty"  db:"-"`
}

type Perfil struct {
	ID              uuid.UUID   `json:"id,omitempty"  db:"id"`
	Nome            string      `json:"nome,omitempty"  db:"nome"`
	ListaPermissoes []Permissao `json:"lista_permissoes,omitempty"  db:"-"`
}

type Permissao struct {
	ID    uuid.UUID `json:"id,omitempty"  db:"id"`
	Nome  string    `json:"nome,omitempty"  db:"nome"`
	Chave string    `json:"chave,omitempty"  db:"chave"`
}

type Token struct {
	UUID         string    `json:"uuid,omitempty" db:"uuid"`
	Usuario_uuid string    `json:"usuario_uuid,omitempty"  db:"usuario_uuid"`
	Perfis       string    `json:"perfis,omitempty"  db:"perfis"`
	DataCriacao  time.Time `json:"data_criacao,omitempty" db:"data_criacao"`
	Validade     string    `json:"validade,omitempty"  db:"validade"`
	Comentarios  string    `json:"comentarios,omitempty"  db:"comentarios"`
}
