package entidades

import (
	"time"

	"github.com/google/uuid"
)

type Usuario struct {
	UUID        uuid.UUID
	Nome        string
	Email       string
	Senha       string
	Perfis      []string
	DataCriacao time.Time
}

// Um Cliente "tem um" Usuario, mas em Go, os campos de Usuario são promovidos para o Cliente.
type Cliente struct {
	Usuario
	CPF        string
	Endereco   string
	TipoPessoa string
}

// Colaborador incorpora Usuario.
type Colaborador struct {
	Usuario
	Cargo string
}
