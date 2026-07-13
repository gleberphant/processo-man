package manipuladores

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

// servindo como interface entre a camada de apresentação e os casos de uso.
type AreaSecretaria struct {
	cduProcesso ICDUProcesso
	cduUsuario  ICDUUsuario
}

// NovoManipuladorAreaSecretaria cria e retorna uma nova instância de ManipuladorAreaSecretaria.
func NovoManipuladorAreaSecretaria(CasosDeUsoProcesso ICDUProcesso, CasosDeUsoUsuario ICDUUsuario) *AreaSecretaria {
	return &AreaSecretaria{
		cduProcesso: CasosDeUsoProcesso,
		cduUsuario:  CasosDeUsoUsuario,
	}
}

// area do cliente
func (m *AreaSecretaria) AreaSecretariaPageListarProcessos(w http.ResponseWriter, r *http.Request) {

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
func (m *AreaSecretaria) AreaSecretariaPageVerProcesso(w http.ResponseWriter, r *http.Request) {

	viewModel := struct {
		Processos []entidades.Processo
	}{
		Processos: []entidades.Processo{},
	}

	apresentacao.ExibirPaginaHTML("area_cliente/page-ver-processo.html", w, r, viewModel)
}
