package entidades

import (
	"github.com/google/uuid"
)

type Permissao struct {
	UUID  uuid.UUID `json:"id,omitempty"  db:"id"`
	Nome  string    `json:"nome,omitempty"  db:"nome"`
	Chave string    `json:"chave,omitempty"  db:"chave"`
}
