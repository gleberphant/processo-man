package area_cliente

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
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

// servindo como interface entre a camada de apresentação e os casos de uso.
type ManipuladorAreaCliente struct {
	cduProcesso ICDUProcesso
	cduUsuario  ICDUUsuario
}

// NovoManipuladorAreaCliente cria e retorna uma nova instância de ManipuladorAreaCliente.
func NovoManipuladorAreaCliente(CasosDeUsoProcesso ICDUProcesso, CasosDeUsoUsuario ICDUUsuario) *ManipuladorAreaCliente {
	return &ManipuladorAreaCliente{
		cduProcesso: CasosDeUsoProcesso,
		cduUsuario:  CasosDeUsoUsuario,
	}
}

// area do cliente
func (m *ManipuladorAreaCliente) AreaClientePageListarProcessos(w http.ResponseWriter, r *http.Request) {

	cliente_uuid, err := uuid.Parse(r.PathValue("cliente_uuid"))

	if err != nil {
		apresentacao.ExibirErro(w, "Cliente inválido")
	}

	listaProcesso, err := m.cduProcesso.ListarProcessosPorCliente(cliente_uuid)

	viewModel := struct {
		Processos []entidades.Processo
	}{
		Processos: listaProcesso,
	}

	apresentacao.ExibirPaginaHTML("area_cliente/page-listar-processos.html", w, viewModel)
}

// area do cliente
func (m *ManipuladorAreaCliente) AreaClientePageVerProcesso(w http.ResponseWriter, r *http.Request) {

	viewModel := struct {
		Processos []entidades.Processo
	}{
		Processos: []entidades.Processo{},
	}

	apresentacao.ExibirPaginaHTML("area_cliente/page-ver-processo.html", w, viewModel)
}
