/*
paginas da área do cliente na qual o cliente poderá
- listar os seus processos
- ver detalhes do seu processo
- ver detalhes do seu cadastro
- inserir documentos no seu cadastro
*/

package manipuladores

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

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

// LISTAR OS PROCESSOS DE UM CLIENTE
func (m *ManipuladorAreaCliente) AreaClientePageListarMeusProcessos(w http.ResponseWriter, r *http.Request) {

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

	apresentacao.ExibirPaginaHTML("area_cliente/page-listar-processos.html", w, r, viewModel)
}

// area do cliente
func (m *ManipuladorAreaCliente) AreaClientePageVerProcesso(w http.ResponseWriter, r *http.Request) {

	viewModel := struct {
		Processos []entidades.Processo
	}{
		Processos: []entidades.Processo{},
	}

	apresentacao.ExibirPaginaHTML("area_cliente/page-ver-processo.html", w, r, viewModel)
}
