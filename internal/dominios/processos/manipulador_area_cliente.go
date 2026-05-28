package processos

import (
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

// area do cliente
func (m *ManipuladorProcesso) AreaClienteListarClientes(w http.ResponseWriter, r *http.Request) {

	lista, err := m.cduUsuario.ListarClientes()
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao listar clientes: %v", err))
		return
	}

	viewModel := ViewModelProcesso{
		Usuarios: lista,
	}

	apresentacao.ExibirPaginaHTML("usuario/page-listar-cliente.html", w, viewModel)
}

func (m *ManipuladorProcesso) AreaClientePageListarProcessos(w http.ResponseWriter, r *http.Request) {

	cliente_uuid, err := uuid.Parse(r.PathValue("cliente_uuid"))

	if err != nil {
		apresentacao.ExibirErro(w, "Cliente inválido")
	}

	listaProcesso, err := m.cduProcesso.ListarProcessosPorCliente(cliente_uuid)

	viewModel := ViewModelProcesso{
		Processos: listaProcesso,
	}

	apresentacao.ExibirPaginaHTML("area_cliente/page-listar-processos.html", w, viewModel)
}

// area do cliente
func (m *ManipuladorProcesso) AreaClientePageVerProcesso(w http.ResponseWriter, r *http.Request) {

	viewModel := ViewModelProcesso{
		Processos: entidades.Processo{},
	}

	apresentacao.ExibirPaginaHTML("area_cliente/page-ver-processo.html", w, viewModel)
}
