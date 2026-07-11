package manipuladores

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

type ICDUUsuario interface {
	ListarClientes() ([]entidades.Cliente, error)
}

type ICDUProcesso interface {
	ListarProcessosPorCliente(uuid.UUID) ([]entidades.Processo, error)
}

type ViewModel struct {
}
