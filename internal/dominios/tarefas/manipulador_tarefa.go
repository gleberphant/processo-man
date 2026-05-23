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

type ViewModelTarefa struct {
	UUID            string
	ProcessoUUID    string
	ResponsavelUUID string
	Tarefa          interface{}
}

// NovoManipuladorTarefa cria e retorna uma nova instância de ManipuladorTarefa.
func NovoManipuladorTarefa(CasosDeUsoTarefa *CDUTarefa) *ManipuladorTarefa {
	return &ManipuladorTarefa{
		cduTarefa: CasosDeUsoTarefa,
	}
}

// PageCriar renderiza o formulário para criação de um novo Tarefa.
func (m *ManipuladorTarefa) PageCriarTarefa(w http.ResponseWriter, r *http.Request) {

	strProcessoUUID := r.URL.Query().Get("ProcessoUUID")

	ProcessoUUID, err := m.cduTarefa.AutenticarProcesso(strProcessoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Criar Tarefa:%v", err))
		return
	}

	viewModel := ViewModelTarefa{
		ProcessoUUID: strProcessoUUID,
		Tarefa:       Tarefa{ProcessoUUID: ProcessoUUID},
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-criar-tarefa.html", w, viewModel)
}

// PageListar renderiza a página contendo a listagem de todos os Tarefas.
func (m *ManipuladorTarefa) PageListarTarefasPorProcesso(w http.ResponseWriter, r *http.Request) {

	strProcessoUUID := r.URL.Query().Get("ProcessoUUID")

	lista, err := m.cduTarefa.ListarTarefasPorProcesso(strProcessoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Listar Tarefa:%v", err))
		return
	}

	viewModel := ViewModelTarefa{
		ProcessoUUID: strProcessoUUID,
		Tarefa:       lista,
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-listar-tarefas.html", w, viewModel)

}

// PageListar renderiza a página contendo a listagem de todos os Tarefas.
func (m *ManipuladorTarefa) PageListarTarefas(w http.ResponseWriter, r *http.Request) {

	lista, err := m.cduTarefa.ListarTarefas()

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Listar Tarefa:%v", err))
		return
	}

	viewModel := ViewModelTarefa{
		Tarefa: lista,
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-listar-tarefas.html", w, viewModel)

}

func (m *ManipuladorTarefa) PageVerTarefa(w http.ResponseWriter, r *http.Request) {
	strUUID := r.PathValue("UUID")

	tarefa, err := m.cduTarefa.BuscarTarefaPorUUID(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("erro PageVerTarefa:%v", err))
		return
	}

	viewModel := ViewModelTarefa{
		UUID:   strUUID,
		Tarefa: tarefa,
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-ver-tarefa.html", w, viewModel)
}

// PageEditar carrega os dados de um Tarefa existente e renderiza o mesmo formulário.
func (m *ManipuladorTarefa) PageEditarTarefa(w http.ResponseWriter, r *http.Request) {
	strUUID := r.PathValue("UUID")

	tarefa, err := m.cduTarefa.BuscarTarefaPorUUID(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Editar Tarefa:%v", err))
		return
	}

	viewModel := ViewModelTarefa{
		UUID:   strUUID,
		Tarefa: *tarefa,
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-criar-tarefa.html", w, viewModel)
}

// --------
func (m *ManipuladorTarefa) CriarTarefaPost(w http.ResponseWriter, r *http.Request) {

	//strUUID := r.PathValue("UUID")
	strProcessoUUID := r.PostFormValue("ProcessoUUID")
	//strResponsavelUUID := r.PostFormValue("ResponsavelUUID")

	processoUUID, err := uuid.Parse(strProcessoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Criar Tarefa:%v", err))
		return
	}

	tarefa := Tarefa{
		ProcessoUUID: processoUUID,
		Nome:         r.PostFormValue("Nome"),
		Comentarios:  r.PostFormValue("Comentarios"),
	}

	err = m.cduTarefa.CriarTarefa(tarefa)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Criar Tarefa: %v", err))
		return
	}

	http.Redirect(w, r, "/processos/"+strProcessoUUID, http.StatusSeeOther)

}

func (m *ManipuladorTarefa) EditarTarefaPost(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")
	strProcessoUUID := r.PostFormValue("ProcessoUUID")
	strResponsavelUUID := r.PostFormValue("ResponsavelUUID")

	UUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Tarefa: %v", err))
		return
	}

	processoUUID, err := uuid.Parse(strProcessoUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Tarefa: %v", err))
		return
	}

	responsavelUUID, err := uuid.Parse(strResponsavelUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Tarefa: %v", err))
		return
	}

	var tarefa = Tarefa{
		UUID:            UUID,
		ProcessoUUID:    processoUUID,
		ResponsavelUUID: responsavelUUID,
		Nome:            r.PostFormValue("Nome"),
		Comentarios:     r.PostFormValue("Comentarios"),
	}

	err = m.cduTarefa.EditarTarefa(tarefa)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Tarefa:%v", err))
		return
	}

	http.Redirect(w, r, "/processos/"+strProcessoUUID, http.StatusSeeOther)
}

func (m *ManipuladorTarefa) DeletarTarefaPost(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	err := m.cduTarefa.DeletarTarefa(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Deletar Tarefa: %v", err))
		return
	}

	paginaAnterior := r.Header.Get("Referer")

	if paginaAnterior == "" {
		paginaAnterior = "/tarefas"
	}
	log.Printf("pagina anterior %s", paginaAnterior)

	http.Redirect(w, r, paginaAnterior, http.StatusSeeOther)
}
