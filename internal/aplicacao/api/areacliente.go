package api

/*
paginas da área do cliente na qual o cliente poderá
- listar os seus processos
- ver detalhes do seu processo
- ver detalhes do seu cadastro
- inserir documentos no seu cadastro
*/

import (
	"encoding/json"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/gleberphant/ProcessoMan/internal/dominio/servicos"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

type ApiAreaCliente struct {
	servicoProcesso *servicos.ServicoProcesso
	servicoUsuario  *servicos.ServicoUsuario
}

// NovoApiAreaCliente cria e retorna uma nova instância de ApiAreaCliente.
func NovoApiAreaCliente(servicoProcesso *servicos.ServicoProcesso, servicoUsuario *servicos.ServicoUsuario) *ApiAreaCliente {
	return &ApiAreaCliente{
		servicoProcesso: servicoProcesso,
		servicoUsuario:  servicoUsuario,
	}
}

// LISTAR OS PROCESSOS DE UM CLIENTE
func (m *ApiAreaCliente) PageListarMeusProcessos(w http.ResponseWriter, r *http.Request) {

	cliente_uuid, err := uuid.Parse(r.PathValue("cliente_uuid"))
	if err != nil {
		apresentacao.ExibirErro(w, "Cliente inválido")
		return
	}

	listaProcesso, err := m.servicoProcesso.ListarProcessosPorUsuario(cliente_uuid)
	if err != nil {
		apresentacao.ExibirErro(w, "Erro ao buscar processos")
		return
	}

	payload := struct {
		Processos []entidades.Processo
	}{
		Processos: listaProcesso,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(payload)
	if err != nil {
		apresentacao.ExibirErro(w, "Erro ao buscar processos")
		return
	}

}

// VER DETALHES DE UM PROCESSO
func (m *ApiAreaCliente) AreaClientePageVerProcesso(w http.ResponseWriter, r *http.Request) {

	processo_uuid, err := uuid.Parse(r.PathValue("processo_uuid"))

	if err == nil {
		return
	}

	processo, err := m.servicoProcesso.BuscarProcessoPorUUID(processo_uuid)

	if err == nil {
		//tratar error
		return
	}

	payload := struct {
		Processos entidades.Processo
	}{
		Processos: *processo,
	}

	err = json.NewEncoder(w).Encode(payload)

	if err != nil {
		//tratar error
		return
	}
}
