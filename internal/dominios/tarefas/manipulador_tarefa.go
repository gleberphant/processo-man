package tarefas

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

// ManipuladorTarefa gerencia as requisições HTTP relacionadas ao domínio de Tarefas,
// servindo como interface entre a camada de apresentação e os casos de uso.
type ManipuladorTarefa struct {
	cduTarefa  *CDUTarefa
	cduUsuario *usuarios.CDUUsuario
}

type ViewModelTarefa struct {
	UUID            string
	ProcessoUUID    string
	ResponsavelUUID string
	Tarefa          interface{}
	Usuarios        interface{}
}

// usuarioDTO é privado ao pacote tarefas e define os campos exibidos na seleção.
type usuarioDTO struct {
	UUID string
	Nome string
}

// NovoManipuladorTarefa cria e retorna uma nova instância de ManipuladorTarefa.
func NovoManipuladorTarefa(CasosDeUsoTarefa *CDUTarefa, CasosDeUsoUsuario *usuarios.CDUUsuario) *ManipuladorTarefa {
	return &ManipuladorTarefa{
		cduTarefa:  CasosDeUsoTarefa,
		cduUsuario: CasosDeUsoUsuario,
	}
}

// obterListaUsuariosDTO centraliza a busca de usuários e conversão para DTO.
func (m *ManipuladorTarefa) obterListaUsuariosDTO() ([]usuarioDTO, error) {
	lista, err := m.cduUsuario.ListarUsuarios()
	if err != nil {
		return nil, err
	}

	var listaUsuarioDTO []usuarioDTO
	for _, item := range lista {
		listaUsuarioDTO = append(listaUsuarioDTO, usuarioDTO{UUID: item.UUID.String(), Nome: item.Nome})
	}

	return listaUsuarioDTO, nil
}

// PageCriar renderiza o formulário para criação de um novo Tarefa.
func (m *ManipuladorTarefa) PageCriarTarefa(w http.ResponseWriter, r *http.Request) {

	strProcessoUUID := r.URL.Query().Get("ProcessoUUID")

	err := m.cduTarefa.ValidarProcesso(strProcessoUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Criar Tarefa:%v", err))
		return
	}

	listaUsuarioDTO, err := m.obterListaUsuariosDTO()
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao carregar usuarios: %v", err))
		return
	}

	processoUUID, _ := uuid.Parse(strProcessoUUID)
	viewModel := ViewModelTarefa{
		ProcessoUUID: strProcessoUUID,
		Tarefa:       Tarefa{ProcessoUUID: processoUUID},
		Usuarios:     listaUsuarioDTO,
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-criar-tarefa.html", w, viewModel)
}

// PageListar renderiza a página contendo a listagem de todos os Tarefas.
func (m *ManipuladorTarefa) PageListarTarefasPorProcesso(w http.ResponseWriter, r *http.Request) {

	strProcessoUUID := r.PathValue("UUID") //r.URL.Query().Get("ProcessoUUID")

	lista, err := m.cduTarefa.ListarTarefasPorProcesso(strProcessoUUID)

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

	listaUsuarioDTO, err := m.obterListaUsuariosDTO()
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao carregar usuarios: %v", err))
		return
	}

	viewModel := ViewModelTarefa{
		UUID:     strUUID,
		Tarefa:   *tarefa,
		Usuarios: listaUsuarioDTO,
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

	tarefa := Tarefa{
		ProcessoUUID:    processoUUID,
		Nome:            r.PostFormValue("Nome"),
		ResponsavelUUID: responsavelUUID,
		Comentarios:     r.PostFormValue("Comentarios"),
	}

	err = m.cduTarefa.CriarTarefa(tarefa)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Criar Tarefa: %v", err))
		return
	}

	http.Redirect(w, r, "/processos/"+processoUUID.String(), http.StatusSeeOther)

}

func (m *ManipuladorTarefa) EditarTarefaPost(w http.ResponseWriter, r *http.Request) {

	UUID, err := uuid.Parse(r.PathValue("UUID"))
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Tarefa: %v", err))
		return
	}

	processoUUID, err := uuid.Parse(r.PostFormValue("ProcessoUUID"))
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Tarefa: %v", err))
		return
	}

	responsavelUUID, err := uuid.Parse(r.PostFormValue("ResponsavelUUID"))

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

	http.Redirect(w, r, "/processos/"+processoUUID.String(), http.StatusSeeOther)
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
