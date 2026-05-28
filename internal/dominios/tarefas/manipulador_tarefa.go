package tarefas

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

// ManipuladorTarefa gerencia as requisições HTTP relacionadas ao domínio de Tarefas,
// servindo como interface entre a camada de apresentação e os casos de uso.
type ICDUUsuario interface {
	Fechar() error
	ListarUsuarios() ([]entidades.Usuario, error)
}

type ManipuladorTarefa struct {
	cduTarefa  *CDUTarefa
	cduUsuario ICDUUsuario
}

// NovoManipuladorTarefa cria e retorna uma nova instância de ManipuladorTarefa.
func NovoManipuladorTarefa(CasosDeUsoTarefa *CDUTarefa, CasosDeUsoUsuario ICDUUsuario) *ManipuladorTarefa {
	return &ManipuladorTarefa{
		cduTarefa:  CasosDeUsoTarefa,
		cduUsuario: CasosDeUsoUsuario,
	}
}

func (m *ManipuladorTarefa) Fechar() {
	m.cduTarefa.Fechar()
	m.cduUsuario.Fechar()
}

// obterListaUsuariosView centraliza a busca de usuários e conversão para DTO.
func (m *ManipuladorTarefa) obterListaUsuariosView() ([]usuarioView, error) {

	lista, err := m.cduUsuario.ListarUsuarios()

	if err != nil {
		return nil, err
	}

	var listaUsuarioView []usuarioView
	for _, item := range lista {
		listaUsuarioView = append(listaUsuarioView, usuarioView{UUID: item.UUID.String(), Nome: item.Nome})
	}

	return listaUsuarioView, nil
}

// PageCriar renderiza o formulário para criação de um novo Tarefa.
func (m *ManipuladorTarefa) PageCriarTarefa(w http.ResponseWriter, r *http.Request) {

	strProcessoUUID := r.URL.Query().Get("ProcessoUUID")

	processoUUID, err := uuid.Parse(strProcessoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do processo inválido: %v", err))
		return
	}

	err = m.cduTarefa.ValidarProcesso(processoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Criar Tarefa:%v", err))
		return
	}

	listaUsuarioView, err := m.obterListaUsuariosView()

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao carregar usuarios: %v", err))
		return
	}

	viewModel := ViewModelTarefa{
		ProcessoUUID: strProcessoUUID,
		Tarefa:       tarefaView{ProcessoUUID: processoUUID},
		Usuarios:     listaUsuarioView,
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-criar-tarefa.html", w, viewModel)
}

// PageListar renderiza a página contendo a listagem de todos os Tarefas.
func (m *ManipuladorTarefa) PageListarTarefasPorProcesso(w http.ResponseWriter, r *http.Request) {

	strProcessoUUID := r.PathValue("processo_uuid") //r.URL.Query().Get("ProcessoUUID")

	processoUUID, err := uuid.Parse(strProcessoUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do processo inválido: %v", err))
		return
	}

	lista, err := m.cduTarefa.ListarTarefasPorProcesso(processoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Listar Tarefa: %v", err))
		return
	}

	viewModel := ViewModelTarefa{
		ProcessoUUID: strProcessoUUID,
		Tarefa:       lista,
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-listar-tarefas.html", w, viewModel)

}

// PageListarTarefasPorResponsavel renderiza a página contendo a listagem de tarefas de um responsável específico.
func (m *ManipuladorTarefa) PageListarTarefasPorResponsavel(w http.ResponseWriter, r *http.Request) {

	strResponsavelUUID := r.PathValue("colaborador_uuid")

	responsavelUUID, err := uuid.Parse(strResponsavelUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do responsável inválido: %v", err))
		return
	}

	lista, err := m.cduTarefa.ListarTarefasPorResponsavel(responsavelUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Listar Tarefas Responsável: %v", err))
		return
	}

	viewModel := ViewModelTarefa{
		Tarefa: lista,
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

	tarefaUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID da tarefa inválido: %v", err))
		return
	}

	tarefa, err := m.cduTarefa.BuscarTarefaPorUUID(tarefaUUID)

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

	tarefaUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID da tarefa inválido: %v", err))
		return
	}

	tarefa, err := m.cduTarefa.BuscarTarefaPorUUID(tarefaUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Editar Tarefa:%v", err))
		return
	}

	listaUsuarioView, err := m.obterListaUsuariosView()
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao carregar usuarios: %v", err))
		return
	}

	viewModel := ViewModelTarefa{
		UUID:     strUUID,
		Tarefa:   *tarefa,
		Usuarios: listaUsuarioView,
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-criar-tarefa.html", w, viewModel)
}

// --------
func (m *ManipuladorTarefa) CriarTarefaPost(w http.ResponseWriter, r *http.Request) {

	processoUUID, err := uuid.Parse(r.PostFormValue("ProcessoUUID"))

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Criar Tarefa: %v", err))
		return
	}

	responsavelUUID, err := uuid.Parse(r.PostFormValue("ResponsavelUUID"))

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Criar Tarefa:%v", err))
		return
	}

	tarefa := entidades.Tarefa{
		ProcessoUUID:    processoUUID,
		ResponsavelUUID: responsavelUUID,
		Nome:            r.PostFormValue("Nome"),
		Comentarios:     r.PostFormValue("Comentarios"),
	}

	err = m.cduTarefa.CriarTarefa(tarefa)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Criar Tarefa: %v", err))
		return
	}

	apresentacao.RedirecionarPaginaAnterior(w, r)

}

func (m *ManipuladorTarefa) EditarTarefaPost(w http.ResponseWriter, r *http.Request) {

	dtoTarefaUUID, err1 := uuid.Parse(r.PathValue("UUID"))
	dtoProcessoUUID, err2 := uuid.Parse(r.PostFormValue("ProcessoUUID"))
	dtoResponsavelUUID, err3 := uuid.Parse(r.PostFormValue("ResponsavelUUID"))

	if err := errors.Join(err1, err2, err3); err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao processar os UUIDs da tarefa: %v", err))
		return
	}

	tarefa := entidades.Tarefa{
		UUID:            dtoTarefaUUID,
		ProcessoUUID:    dtoProcessoUUID,
		ResponsavelUUID: dtoResponsavelUUID,
		Nome:            r.PostFormValue("Nome"),
		Comentarios:     r.PostFormValue("Comentarios"),
	}

	if err := m.cduTarefa.EditarTarefa(tarefa); err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Tarefa:%v", err))
		return
	}
	apresentacao.RedirecionarPaginaAnterior(w, r)

}

func (m *ManipuladorTarefa) DeletarTarefaPost(w http.ResponseWriter, r *http.Request) {

	tarefaUUID, err := uuid.Parse(r.PathValue("UUID"))

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID da tarefa inválido: %v", err))
		return
	}

	err = m.cduTarefa.DeletarTarefa(tarefaUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Deletar Tarefa: %v", err))
		return
	}

	apresentacao.RedirecionarPaginaAnterior(w, r)
}
