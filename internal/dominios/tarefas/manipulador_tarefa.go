package tarefas

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

// ManipuladorTarefa gerencia as requisições HTTP relacionadas ao domínio de Tarefas,
// servindo como interface entre a camada de apresentação e os casos de uso.
type ManipuladorTarefa struct {
	cduTarefa *CDUTarefa
}

// NovoManipuladorTarefa cria e retorna uma nova instância de ManipuladorTarefa.
func NovoManipuladorTarefa(CasosDeUsoTarefa *CDUTarefa) *ManipuladorTarefa {
	return &ManipuladorTarefa{
		cduTarefa: CasosDeUsoTarefa,
	}
}

// PageCriar renderiza o formulário para criação de um novo Tarefa.
func (m *ManipuladorTarefa) PageCriarTarefa(w http.ResponseWriter, r *http.Request) {

	viewModel := struct {
		UUID            string
		ProcessoUUID    string
		Nome            string
		ResponsavelUUID string
		Comentarios     string
	}{
		ProcessoUUID: r.URL.Query().Get("ProcessoUUID"),
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-criar-tarefa.html", w, viewModel)
}

// PageListar renderiza a página contendo a listagem de todos os Tarefas.
func (m *ManipuladorTarefa) PageListarTarefas(w http.ResponseWriter, r *http.Request) {

	strProcessoUUID := r.URL.Query().Get("ProcessoUUID")

	lista, err := m.cduTarefa.ListarTarefas(strProcessoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Listar Tarefa:%v", err))
		return
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-listar-tarefas.html", w, lista)

}

// PageEditar carrega os dados de um Tarefa existente e renderiza o mesmo formulário.
func (m *ManipuladorTarefa) PageEditarTarefa(w http.ResponseWriter, r *http.Request) {
	uuidStr := r.URL.Query().Get("uuid")

	tarefa, err := m.cduTarefa.BuscarTarefaPorUUID(uuidStr)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Editar Tarefa:%v", err))
		return
	}

	viewModel := struct {
		UUID            string
		ProcessoUUID    string
		Nome            string
		ResponsavelUUID string
		Comentarios     string
	}{
		UUID:            tarefa.UUID.String(),
		ProcessoUUID:    tarefa.ProcessoUUID.String(),
		Nome:            tarefa.Nome,
		ResponsavelUUID: tarefa.ResponsavelUUID.String(),
		Comentarios:     tarefa.Comentarios,
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-criar-Tarefa.html", w, viewModel)
}

// --------
func (m *ManipuladorTarefa) CriarTarefaPost(w http.ResponseWriter, r *http.Request) {
	processoUUID, err := uuid.Parse(r.PostFormValue("ProcessoUUID"))
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Criar Tarefa:%v", err))
		return
	}

	//UUID, err := uuid.Parse(r.PostFormValue("uuid"))
	tarefa := Tarefa{
		ProcessoUUID: processoUUID,
		Nome:         r.PostFormValue("Nome"),
		Comentarios:  r.PostFormValue("Comentarios"),
	}

	log.Printf("tarefa recebida %v", tarefa)

	err = m.cduTarefa.CriarTarefa(tarefa)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Criar Tarefa:%v", err))
		return
	}

	http.Redirect(w, r, "/processo/visualizar?uuid="+tarefa.ProcessoUUID.String(), http.StatusSeeOther)

}

func (m *ManipuladorTarefa) EditarTarefaPost(w http.ResponseWriter, r *http.Request) {

	UUID, err := uuid.Parse(r.PostFormValue("uuid"))
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Tarefa: %v", err))
		return
	}

	uuid.MustParse(r.PostFormValue("ProcessoUUID"))

	var tarefa = Tarefa{
		UUID:            UUID,
		ProcessoUUID:    uuid.MustParse(r.PostFormValue("ProcessoUUID")),
		ResponsavelUUID: uuid.MustParse(r.PostFormValue("ResponsavelUUID")),
		Nome:            r.PostFormValue("Nome"),
		Comentarios:     r.PostFormValue("Comentarios"),
	}

	err = m.cduTarefa.EditarTarefa(tarefa)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Tarefa:%v", err))
		return
	}

	http.Redirect(w, r, "/tarefa/listar", http.StatusSeeOther)
}

func (m *ManipuladorTarefa) DeletarTarefaPost(w http.ResponseWriter, r *http.Request) {

	var UUID = r.PostFormValue("uuid")

	err := m.cduTarefa.DeletarTarefa(UUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Tarefa:%v", err))
		return
	}

	http.Redirect(w, r, "/Tarefa/listar", http.StatusSeeOther)
}
