package processos

import (
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

// servindo como interface entre a camada de apresentação e os casos de uso.
type ManipuladorProcesso struct {
	cduProcesso *CDUProcesso
}

type ViewModelProcesso struct {
	UUID     string
	Processo interface{}
	Anexos   []string
}

// NovoManipuladorProcesso cria e retorna uma nova instância de ManipuladorProcesso.
func NovoManipuladorProcesso(CasosDeUsoProcesso *CDUProcesso) *ManipuladorProcesso {
	return &ManipuladorProcesso{
		cduProcesso: CasosDeUsoProcesso,
	}
}

// PageCriar renderiza o formulário para criação de um novo processo.
func (m *ManipuladorProcesso) PageCriar(w http.ResponseWriter, r *http.Request) {

	viewModel := ViewModelProcesso{
		Processo: Processo{},
	}

	apresentacao.ExibirPaginaHTML("processo/page-criar-processo.html", w, viewModel)
}

// PageListar renderiza a página contendo a listagem de todos os processos.
func (m *ManipuladorProcesso) PageListar(w http.ResponseWriter, r *http.Request) {

	lista, err := m.cduProcesso.ListarProcessos()

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("erro Page Listar Processo:%v", err))
	}

	viewModel := ViewModelProcesso{
		Processo: lista,
	}

	apresentacao.ExibirPaginaHTML("processo/page-listar-processos.html", w, viewModel)

}

// PageEditar carrega os dados de um processo existente e renderiza o mesmo formulário.
func (m *ManipuladorProcesso) PageVerProcesso(w http.ResponseWriter, r *http.Request) {
	strUUID := r.PathValue("UUID")

	processo, err := m.cduProcesso.BuscarProcessoPorUUID(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("erro Page Ver Processo:%v", err))
		return
	}

	viewModel := ViewModelProcesso{
		UUID:     strUUID,
		Processo: *processo,
		Anexos:   []string{"arquivo1.doc", "arquivo2.doc"},
	}

	apresentacao.ExibirPaginaHTML("processo/page-ver-processo.html", w, viewModel)
}

// PageEditar carrega os dados de um processo existente e renderiza o mesmo formulário.
func (m *ManipuladorProcesso) PageEditar(w http.ResponseWriter, r *http.Request) {
	strUUID := r.PathValue("UUID")

	processo, err := m.cduProcesso.BuscarProcessoPorUUID(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro PageEditar:%v", err))
		return
	}

	viewModel := ViewModelProcesso{
		UUID:     strUUID,
		Processo: processo,
		//Anexos:   []string{"arquivo1.doc", "arquivo2.doc"},
	}

	apresentacao.ExibirPaginaHTML("processo/page-criar-processo.html", w, viewModel)
}

// CriarProcessoPost processa a submissão do formulário para persistir um novo processo.
func (m *ManipuladorProcesso) CriarProcessoPost(w http.ResponseWriter, r *http.Request) {

	var Processo = Processo{
		Nome: r.PostFormValue("nome"),
	}

	err := m.cduProcesso.CriarProcesso(Processo)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro criar Processo: %v", err))
	}

	http.Redirect(w, r, "/processos", http.StatusSeeOther)

}

// EditarProcessoPost processa a atualização de um processo existente.
func (m *ManipuladorProcesso) EditarProcessoPost(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	UUID, err := uuid.Parse(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro editar Processo:%v", err))
	}

	processoDTO := Processo{
		UUID: UUID,
		Nome: r.PostFormValue("nome"),
	}

	err = m.cduProcesso.EditarProcesso(processoDTO)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro editar Processo:%v", err))
	}

	http.Redirect(w, r, "/processos", http.StatusSeeOther)
}

// DeletarProcessoPost remove um processo com base no identificador enviado via formulário.
func (m *ManipuladorProcesso) DeletarProcessoPost(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	err := m.cduProcesso.DeletarProcesso(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro deletar Processo: %v", err))
	}

	http.Redirect(w, r, "/processos", http.StatusSeeOther)
}
