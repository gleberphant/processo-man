package autenticacao

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	UUID        uuid.UUID `json:"uuid,omitempty" db:"uuid"`
	UsuarioUUID uuid.UUID `json:"usuario_uuid,omitempty"  db:"usuario_uuid"`
	Perfis      string    `json:"perfis,omitempty"  db:"perfis"`
	DataCriacao time.Time `json:"data_criacao,omitempty" db:"data_criacao"`
	Validade    string    `json:"validade,omitempty"  db:"validade"`
	Comentarios string    `json:"comentarios,omitempty"  db:"comentarios"`
}
