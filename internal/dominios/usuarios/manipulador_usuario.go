package usuarios

import (
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

type ICDUTarefas interface {
	Fechar() error
	ListarTarefasPorResponsavel(responsavelUUID uuid.UUID) ([]entidades.Tarefa, error)
}

type ManipuladorUsuario struct {
	cduUsario *CDUUsuario
	cduTarefa ICDUTarefas
}

func NovoManipuladorUsuario(casosDeUsoUsuario *CDUUsuario, casosDeUsoTarefa ICDUTarefas) *ManipuladorUsuario {

	return &ManipuladorUsuario{
		cduUsario: casosDeUsoUsuario,
		cduTarefa: casosDeUsoTarefa,
	}
}

func (m *ManipuladorUsuario) Fechar() {
	m.cduUsario.Fechar()

}

func (m *ManipuladorUsuario) PageCriarUsuario(w http.ResponseWriter, r *http.Request) {

	viewModel := ViewModelUsuario{}

	apresentacao.ExibirPaginaHTML("usuario/page-criar-usuario.html", w, viewModel)

}

func (m *ManipuladorUsuario) PageCriarCliente(w http.ResponseWriter, r *http.Request) {
	viewModel := ViewModelUsuario{}

	apresentacao.ExibirPaginaHTML("usuario/page-criar-cliente.html", w, viewModel)
}

func (m *ManipuladorUsuario) PageCriarColaborador(w http.ResponseWriter, r *http.Request) {
	viewModel := ViewModelUsuario{}

	apresentacao.ExibirPaginaHTML("usuario/page-criar-colaborador.html", w, viewModel)
}

func (m *ManipuladorUsuario) PageListarUsuarios(w http.ResponseWriter, r *http.Request) {

	lista, err := m.cduUsario.ListarUsuarios()

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro PageListar:%v", err))
		return
	}

	viewModel := ViewModelUsuario{
		Usuarios: lista,
	}

	apresentacao.ExibirPaginaHTML("usuario/page-listar-usuario.html", w, viewModel)

}

func (m *ManipuladorUsuario) PageListarClientes(w http.ResponseWriter, r *http.Request) {

	lista, err := m.cduUsario.ListarClientes()
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao listar clientes: %v", err))
		return
	}

	viewModel := ViewModelUsuario{
		Usuarios: lista,
	}

	apresentacao.ExibirPaginaHTML("usuario/page-listar-cliente.html", w, viewModel)
}

func (m *ManipuladorUsuario) PageListarColaboradores(w http.ResponseWriter, r *http.Request) {

	lista, err := m.cduUsario.ListarColaboradores()
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao listar colaboradores: %v", err))
		return
	}

	viewModel := ViewModelUsuario{
		Usuarios: lista,
	}

	apresentacao.ExibirPaginaHTML("usuario/page-listar-colaborador.html", w, viewModel)
}

func (m *ManipuladorUsuario) PageEditarUsuario(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	usuarioUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do usuário inválido: %v", err))
		return
	}

	usuario, err := m.cduUsario.BuscarUsuarioPorUUID(usuarioUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro PageEditar:%v", err))
		return
	}

	viewModel := ViewModelUsuario{
		UUID:     strUUID,
		Usuarios: usuario,
	}

	apresentacao.ExibirPaginaHTML("usuario/page-criar-usuario.html", w, viewModel)

}

func (m *ManipuladorUsuario) obterListaTarefasViewModel(usuarioUUID uuid.UUID) ([]tarefasView, error) {

	tarefas, err := m.cduTarefa.ListarTarefasPorResponsavel(usuarioUUID)

	if err != nil {
		return nil, err
	}
	listaTarefasView := []tarefasView{}

	for _, tarefa := range tarefas {
		tarefaView := tarefasView{
			UUID:            tarefa.UUID,
			ProcessoUUID:    tarefa.ProcessoUUID,
			ResponsavelUUID: tarefa.ResponsavelUUID,
			Nome:            tarefa.Nome,
			Concluida:       tarefa.Concluida,
			Comentarios:     tarefa.Comentarios,
			DataConclusao:   tarefa.DataConclusao,
			DataCriacao:     tarefa.DataCriacao,
		}

		listaTarefasView = append(listaTarefasView, tarefaView)
	}

	return listaTarefasView, nil
}

func (m *ManipuladorUsuario) PageVerUsuario(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	usuarioUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do usuário inválido: %v", err))
		return
	}

	usuario, err := m.cduUsario.BuscarUsuarioPorUUID(usuarioUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Editar:%v", err))
		return
	}

	listaTarefasView, err := m.obterListaTarefasViewModel(usuarioUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Editar:%v", err))
		return
	}

	viewModel := ViewModelUsuario{
		UUID:     strUUID,
		Usuarios: usuario,
		Anexos:   []string{"documento1", "document2"},
		Tarefas:  listaTarefasView,
	}

	apresentacao.ExibirPaginaHTML("usuario/page-ver-usuario.html", w, viewModel)

}

func (m *ManipuladorUsuario) CriarUsuarioPost(w http.ResponseWriter, r *http.Request) {

	var usuario = entidades.Usuario{
		Nome:  r.PostFormValue("nome"),
		Email: r.PostFormValue("email"),
		Senha: r.PostFormValue("senha"),
	}

	err := m.cduUsario.CriaUsuario(&usuario)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Criar Usuario:%v", err))
		return
	}

	http.Redirect(w, r, "/usuarios", http.StatusSeeOther)

}

func (m *ManipuladorUsuario) CriarClientePost(w http.ResponseWriter, r *http.Request) {
	cliente := entidades.Cliente{
		Usuario: entidades.Usuario{
			Nome:  r.PostFormValue("nome"),
			Email: r.PostFormValue("email"),
			Senha: r.PostFormValue("senha"),
		},
		CPF:        r.PostFormValue("cpf"),
		Endereco:   r.PostFormValue("endereco"),
		TipoPessoa: r.PostFormValue("tipo_pessoa"),
	}

	err := m.cduUsario.CriarCliente(&cliente)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Criar Cliente: %v", err))
		return
	}

	http.Redirect(w, r, "/usuarios", http.StatusSeeOther)
}

func (m *ManipuladorUsuario) CriarColaboradorPost(w http.ResponseWriter, r *http.Request) {
	colaborador := entidades.Colaborador{
		Usuario: entidades.Usuario{
			Nome:  r.PostFormValue("nome"),
			Email: r.PostFormValue("email"),
			Senha: r.PostFormValue("senha"),
		},
		Cargo: r.PostFormValue("cargo"),
	}

	err := m.cduUsario.CriarColaborador(&colaborador)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Criar Colaborador: %v", err))
		return
	}

	http.Redirect(w, r, "/usuarios", http.StatusSeeOther)
}

func (m *ManipuladorUsuario) EditarUsuarioPost(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PostFormValue("uuid") //r.PathValue("UUID")

	UUID, err := uuid.Parse(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro editar Processo:%v", err))
	}

	var usuario = entidades.Usuario{
		UUID:  UUID,
		Nome:  r.PostFormValue("nome"),
		Email: r.PostFormValue("email"),
		Senha: r.PostFormValue("senha"),
	}

	err = m.cduUsario.EditarUsuario(usuario)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Usuario:%v", err))
	}

	http.Redirect(w, r, "/usuarios", http.StatusSeeOther)
}

func (m *ManipuladorUsuario) DeletarUsuarioPost(w http.ResponseWriter, r *http.Request) {

	var strUUID = r.PathValue("UUID") //r.PostFormValue("uuid")

	usuarioUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do usuário inválido: %v", err))
		return
	}

	err = m.cduUsario.DeletarUsuario(usuarioUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Usuario:%v", err))
		return
	}

	http.Redirect(w, r, "/usuarios", http.StatusSeeOther)
}
