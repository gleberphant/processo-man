package processos

import (
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominios/tarefas"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

// servindo como interface entre a camada de apresentação e os casos de uso.
type ManipuladorProcesso struct {
	cduProcesso *CDUProcesso
}

// NovoManipuladorProcesso cria e retorna uma nova instância de ManipuladorProcesso.
func NovoManipuladorProcesso(CasosDeUsoProcesso *CDUProcesso) *ManipuladorProcesso {
	return &ManipuladorProcesso{
		cduProcesso: CasosDeUsoProcesso,
	}
}

// PageCriar renderiza o formulário para criação de um novo processo.
func (m *ManipuladorProcesso) PageCriar(w http.ResponseWriter, r *http.Request) {

	apresentacao.ExibirPaginaHTML("processo/page-criar-processo.html", w, Processo{})
}

// PageListar renderiza a página contendo a listagem de todos os processos.
func (m *ManipuladorProcesso) PageListar(w http.ResponseWriter, r *http.Request) {

	lista, err := m.cduProcesso.ListarProcessos()

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("erro PageListar:%v", err))
	}

	apresentacao.ExibirPaginaHTML("processo/page-listar-processos.html", w, lista)

}

// PageEditar carrega os dados de um processo existente e renderiza o mesmo formulário.
func (m *ManipuladorProcesso) PageVisualizarProcesso(w http.ResponseWriter, r *http.Request) {
	uuidStr := r.PathValue("uuid")

	processo, err := m.cduProcesso.BuscarProcessoPorUUID(uuidStr)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("erro PageVer:%v", err))
		return
	}

	viewModel := struct {
		Titulo   string
		Mensagem string
		Processo Processo
		Anexos   []string
		Tarefas  []tarefas.Tarefa
	}{
		Titulo:   "Processo nº: " + string(processo.UUID.String()),
		Mensagem: "ok",
		Processo: *processo,
		Anexos:   []string{"arquivo1.doc", "arquivo2.doc"},
		Tarefas:  processo.Tarefas,
	}

	apresentacao.ExibirPaginaHTML("processo/page-ver-processo.html", w, viewModel)
}

// PageEditar carrega os dados de um processo existente e renderiza o mesmo formulário.
func (m *ManipuladorProcesso) PageEditar(w http.ResponseWriter, r *http.Request) {
	uuidStr := r.PathValue("uuid")

	processo, err := m.cduProcesso.BuscarProcessoPorUUID(uuidStr)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro PageEditar:%v", err))
		return
	}

	apresentacao.ExibirPaginaHTML("processo/page-criar-processo.html", w, processo)
}

// PageEditar carrega os dados de um processo existente e renderiza o mesmo formulário.
func (m *ManipuladorProcesso) PageDeletar(w http.ResponseWriter, r *http.Request) {
	uuidStr := r.PathValue("uuid")

	processo, err := m.cduProcesso.BuscarProcessoPorUUID(uuidStr)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("erro PageDeletar:%v", err))
	}

	// Reutiliza o mesmo template, injetando os dados do processo
	apresentacao.ExibirPaginaHTML("processo/page-criar-processo.html", w, processo)
}

// CriarProcessoPost processa a submissão do formulário para persistir um novo processo.
func (m *ManipuladorProcesso) CriarProcessoPost(w http.ResponseWriter, r *http.Request) {

	UUID, err := uuid.Parse(r.PostFormValue("uuid"))
	var Processo = Processo{
		UUID: UUID,
		Nome: r.PostFormValue("nome"),
		Tarefas: []tarefas.Tarefa{{
			Nome: r.PostFormValue("tarefa"),
		},
		},
	}

	err = m.cduProcesso.CriarProcesso(Processo)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro criar Processo:%v", err))
	}

	http.Redirect(w, r, "/processo/listar", http.StatusSeeOther)

}

// EditarProcessoPost processa a atualização de um processo existente.
func (m *ManipuladorProcesso) EditarProcessoPost(w http.ResponseWriter, r *http.Request) {

	UUID, err := uuid.Parse(r.PostFormValue("uuid"))
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro editar Processo:%v", err))
	}

	var ProcessoDTO = Processo{
		UUID: UUID,
		Nome: r.PostFormValue("nome"),
	}

	err = m.cduProcesso.EditarProcesso(ProcessoDTO)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro editar Processo:%v", err))
	}

	http.Redirect(w, r, "/processo/listar", http.StatusSeeOther)
}

// DeletarProcessoPost remove um processo com base no identificador enviado via formulário.
func (m *ManipuladorProcesso) DeletarProcessoPost(w http.ResponseWriter, r *http.Request) {

	var UUID = r.PostFormValue("uuid")

	err := m.cduProcesso.DeletarProcesso(UUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro deletar Processo:%v", err))
	}

	http.Redirect(w, r, "/processo/listar", http.StatusSeeOther)
}
