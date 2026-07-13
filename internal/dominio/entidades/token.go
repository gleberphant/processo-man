package entidades

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	UUID        uuid.UUID
	UsuarioUUID uuid.UUID
	UsuarioNome string
	DataCriacao time.Time
	Validade    string
	Comentarios string
	Perfis      []string
}

type Permissoes struct {
	Rota   string
	Perfis []string
}
