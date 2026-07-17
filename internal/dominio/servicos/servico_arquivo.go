package servicos

import (
	"github.com/gleberphant/ProcessoMan/internal/aplicacao/repositorios"
	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/google/uuid"
)

type ServicoArquivo struct {
	repoArquivo *repositorios.RepositorioArquivo
}

func NovoServicoArquivo(ArquivoRepo *repositorios.RepositorioArquivo) *ServicoArquivo {

	return &ServicoArquivo{
		repoArquivo: ArquivoRepo,
	}
}

func (a *ServicoArquivo) Fechar() error {
	a.repoArquivo.Fechar()
	return nil
}

func (u *ServicoArquivo) ListarProcessos() ([]entidades.Arquivo, error) {

	return u.repoArquivo.Listar()

}

func (u *ServicoArquivo) ListarArquivosPorUsuario(UUID uuid.UUID) ([]entidades.Arquivo, error) {

	return u.repoArquivo.ListarArquivosPorUsuario(UUID)

}
