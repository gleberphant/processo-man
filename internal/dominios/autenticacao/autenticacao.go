package autenticacao

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	UUID        uuid.UUID
	UsuarioUUID uuid.UUID
	DataCriacao time.Time
	Validade    string
	Comentarios string
	Perfis      string
}

type Permissoes struct {
	UUID      uuid.UUID
	Perfil    string
	Permissao string
}
