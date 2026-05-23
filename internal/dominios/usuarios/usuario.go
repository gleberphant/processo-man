package usuarios

import (
	"time"

	"github.com/google/uuid"
)

type Usuario struct {
	UUID        uuid.UUID `json:"uuid,omitempty"  db:"uuid"`
	Nome        string    `json:"nome,omitempty"  db:"nome"`
	Email       string    `json:"email,omitempty"  db:"email"`
	Senha       string    `json:"-"  db:"senha"`
	DataCriacao time.Time `json:"data_criacao,omitempty"  db:"data_criacao"`
	Perfis      []Perfil  `json:"perfis,omitempty"  db:"-"`
}

type Perfil struct {
	UUID            uuid.UUID   `json:"id,omitempty"  db:"id"`
	Nome            string      `json:"nome,omitempty"  db:"nome"`
	ListaPermissoes []Permissao `json:"lista_permissoes,omitempty"  db:"-"`
}

type Permissao struct {
	UUID  uuid.UUID `json:"id,omitempty"  db:"id"`
	Nome  string    `json:"nome,omitempty"  db:"nome"`
	Chave string    `json:"chave,omitempty"  db:"chave"`
}

type Cliente struct {
	Usuario `json:"usuario,omitempty"  db:"usuario"`
}

type Colaborador struct {
	Usuario `json:"usuario,omitempty"  db:"usuario"`
}
