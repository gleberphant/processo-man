package processos

import (
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

type ICDUUsuario interface {
	ListarClientes() ([]entidades.Cliente, error)
}

// servindo como interface entre a camada de apresentação e os casos de uso.
type ManipuladorProcesso struct {
	cduProcesso *CDUProcesso
	cduUsuario  ICDUUsuario
}

// NovoManipuladorProcesso cria e retorna uma nova instância de ManipuladorProcesso.
func NovoManipuladorProcesso(CasosDeUsoProcesso *CDUProcesso, CasosDeUsoUsuario ICDUUsuario) *ManipuladorProcesso {
	return &ManipuladorProcesso{
		cduProcesso: CasosDeUsoProcesso,
		cduUsuario:  CasosDeUsoUsuario,
	}
}

func (m *ManipuladorProcesso) Fechar() {
	m.cduProcesso.Fechar()
}

// PageCriar renderiza o formulário para criação de um novo processo.
func (m *ManipuladorProcesso) PageCriar(w http.ResponseWriter, r *http.Request) {

	viewModel := ViewModelProcesso{
		Processos: entidades.Processo{},
	}

	apresentacao.ExibirPaginaHTML("processo/page-criar-processo.html", w, r, viewModel)
}

// PageListar renderiza a página contendo a listagem de todos os processos.
func (m *ManipuladorProcesso) PageListar(w http.ResponseWriter, r *http.Request) {

	listaProcessos, err := m.cduProcesso.ListarProcessos()

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("erro Page Listar Processo:%v", err))
		return
	}

	viewModel := ViewModelProcesso{
		Processos: listaProcessos,
	}

	apresentacao.ExibirPaginaHTML("processo/page-listar-processos.html", w, r, viewModel)

}

// PageEditar carrega os dados de um processo existente e renderiza o mesmo formulário.
func (m *ManipuladorProcesso) PageVerProcesso(w http.ResponseWriter, r *http.Request) {
	strUUID := r.PathValue("UUID")

	processoUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do processo inválido: %v", err))
		return
	}

	processo, err := m.cduProcesso.BuscarProcessoPorUUID(processoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("erro Page Ver Processo:%v", err))
		return
	}

	viewModel := ViewModelProcesso{
		UUID:     strUUID,
		Processo: *processo,
		Anexos:   []string{"arquivo1.doc", "arquivo2.doc"},
	}

	apresentacao.ExibirPaginaHTML("processo/page-ver-processo.html", w, r, viewModel)
}

// PageEditar carrega os dados de um processo existente e renderiza o mesmo formulário.
func (m *ManipuladorProcesso) PageEditar(w http.ResponseWriter, r *http.Request) {
	strUUID := r.PathValue("UUID")

	processoUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do processo inválido: %v", err))
		return
	}

	processo, err := m.cduProcesso.BuscarProcessoPorUUID(processoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro PageEditar:%v", err))
		return
	}

	viewModel := ViewModelProcesso{
		UUID:      strUUID,
		Processos: processo,
		//Anexos:   []string{"arquivo1.doc", "arquivo2.doc"},
	}

	apresentacao.ExibirPaginaHTML("processo/page-criar-processo.html", w, r, viewModel)
}

// CriarProcessoPost processa a submissão do formulário para persistir um novo processo.
func (m *ManipuladorProcesso) CriarProcessoPost(w http.ResponseWriter, r *http.Request) {

	var Processo = entidades.Processo{
		Nome: r.PostFormValue("nome"),
	}

	err := m.cduProcesso.CriarProcesso(Processo)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro criar Processo: %v", err))
		return
	}

	http.Redirect(w, r, "/processos", http.StatusSeeOther)

}

// EditarProcessoPost processa a atualização de um processo existente.
func (m *ManipuladorProcesso) EditarProcessoPost(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	UUID, err := uuid.Parse(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro editar Processo:%v", err))
		return
	}

	processoDTO := entidades.Processo{
		UUID: UUID,
		Nome: r.PostFormValue("nome"),
	}

	err = m.cduProcesso.EditarProcesso(processoDTO)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro editar Processo:%v", err))
		return
	}

	http.Redirect(w, r, "/processos", http.StatusSeeOther)
}

// DeletarProcessoPost remove um processo com base no identificador enviado via formulário.
func (m *ManipuladorProcesso) DeletarProcessoPost(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	processoUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do processo inválido: %v", err))
		return
	}

	err = m.cduProcesso.DeletarProcesso(processoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro deletar Processo: %v", err))
		return
	}

	http.Redirect(w, r, "/processos", http.StatusSeeOther)
}
