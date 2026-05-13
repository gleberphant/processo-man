package modelos

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	UUID         uuid.UUID `json:"uuid,omitempty" db:"uuid"`
	Token        string    `json:"token,omitempty" db:"token"`
	Usuario_uuid string    `json:"usuario_uuid,omitempty"  db:"usuario_uuid"`
	Perfis       string    `json:"perfis,omitempty"  db:"perfis"`
	DataCriacao  time.Time `json:"data_criacao,omitempty" db:"data_criacao"`
	Ativo        bool      `json:"ativo,omitempty"  db:"ativo"`
	Validade     string    `json:"validade,omitempty"  db:"validade"`
	Comentarios  string    `json:"comentarios,omitempty"  db:"comentarios"`
}
