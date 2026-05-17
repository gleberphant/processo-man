package casosdeuso

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

type IRepositorioProcesso interface {
	Criar(entidades.Processo) error
	Listar() ([]entidades.Processo, error)
	Atualizar(entidades.Processo) error
	Deletar(uuid.UUID) error
	BuscarPorUUID(uuid.UUID) (*entidades.Processo, error)
}

type CDUProcesso struct {
	repoProcesso IRepositorioProcesso
}

func NovoCDUProcesso(ProcessosRepo IRepositorioProcesso) *CDUProcesso {

	return &CDUProcesso{
		repoProcesso: ProcessosRepo,
	}
}
