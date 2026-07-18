package manipuladores

import (
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/gleberphant/ProcessoMan/internal/dominio/servicos"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

type ManipuladorUsuario struct {
	servicoUsuario  *servicos.ServicoUsuario
	servicoProcesso *servicos.ServicoProcesso
	servicoTarefa   *servicos.ServicoTarefa
}

func NovoManipuladorUsuario(casosDeUsoUsuario *servicos.ServicoUsuario, casosDeUsoProcesso *servicos.ServicoProcesso, casosDeUsoTarefa *servicos.ServicoTarefa) *ManipuladorUsuario {

	return &ManipuladorUsuario{
		servicoUsuario:  casosDeUsoUsuario,
		servicoProcesso: casosDeUsoProcesso,
		servicoTarefa:   casosDeUsoTarefa,
	}
}

func (m *ManipuladorUsuario) Fechar() {
	m.servicoUsuario.Fechar()

}

func (m *ManipuladorUsuario) PageCriarUsuario(w http.ResponseWriter, r *http.Request) {

	viewModel := ViewModelPageUsuario{}

	apresentacao.ExibirPaginaHTML("usuario/page-criar-usuario.html", w, r, viewModel)

}

func (m *ManipuladorUsuario) PageCriarCliente(w http.ResponseWriter, r *http.Request) {
	viewModel := ViewModelPageUsuario{}

	apresentacao.ExibirPaginaHTML("usuario/page-criar-cliente.html", w, r, viewModel)
}

func (m *ManipuladorUsuario) PageCriarColaborador(w http.ResponseWriter, r *http.Request) {
	viewModel := ViewModelPageUsuario{}

	apresentacao.ExibirPaginaHTML("usuario/page-criar-colaborador.html", w, r, viewModel)
}

func (m *ManipuladorUsuario) PageListarUsuario(w http.ResponseWriter, r *http.Request) {

	lista, err := m.servicoUsuario.ListarUsuarios()

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro PageListar:%v", err))
		return
	}

	viewModel := ViewModelPageUsuario{
		Usuario: lista,
	}

	apresentacao.ExibirPaginaHTML("usuario/page-listar-usuario.html", w, r, viewModel)

}

func (m *ManipuladorUsuario) PageListarClientes(w http.ResponseWriter, r *http.Request) {

	lista, err := m.servicoUsuario.ListarClientes()
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao listar clientes: %v", err))
		return
	}

	viewModel := ViewModelPageUsuario{
		Usuario: lista,
	}

	apresentacao.ExibirPaginaHTML("usuario/page-listar-cliente.html", w, r, viewModel)
}

func (m *ManipuladorUsuario) PageListarColaboradores(w http.ResponseWriter, r *http.Request) {

	lista, err := m.servicoUsuario.ListarColaboradores()
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao listar colaboradores: %v", err))
		return
	}

	viewModel := ViewModelPageUsuario{
		Usuario: lista,
	}

	apresentacao.ExibirPaginaHTML("usuario/page-listar-colaborador.html", w, r, viewModel)
}

func (m *ManipuladorUsuario) PageEditarUsuario(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	usuarioUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do usuário inválido: %v", err))
		return
	}

	usuario, err := m.servicoUsuario.BuscarUsuarioPorUUID(usuarioUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro PageEditar:%v", err))
		return
	}

	viewModel := ViewModelPageUsuario{
		UUID:    strUUID,
		Usuario: usuario,
	}

	apresentacao.ExibirPaginaHTML("usuario/page-criar-usuario.html", w, r, viewModel)

}

func (m *ManipuladorUsuario) PageVerUsuario(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	// converter uuid do usuario
	usuarioUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID do usuário inválido: %v", err))
		return
	}

	usuario, err := m.servicoUsuario.BuscarUsuarioPorUUID(usuarioUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro PageVerUsuario:%v", err))
		return
	}

	listaProcesso, err := m.servicoProcesso.ListarProcessosPorUsuario(usuarioUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao buscar processos do usuário:%v", err))
		return
	}

	listaTarefa, err := m.servicoTarefa.ListarTarefasPorResponsavel(usuarioUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Editar:%v", err))
		return
	}

	listaArquivo := []entidades.Arquivo{
		{UUID: uuid.New(), Nome: "arquivo1.txt", URL: "file://123"},
		{UUID: uuid.New(), Nome: "arquivo2.txt", URL: "file://123"},
		{UUID: uuid.New(), Nome: "arquivo3.txt", URL: "file://123"},
		{UUID: uuid.New(), Nome: "arquivo4.txt", URL: "file://123"},
	}

	viewModelPageUsuario := ViewModelPageUsuario{
		UUID:     strUUID,
		Usuario:  usuario,
		Arquivo:  listaArquivo,
		Tarefa:   listaTarefa,
		Processo: listaProcesso,
	}

	apresentacao.ExibirPaginaHTML("usuario/page-ver-usuario.html", w, r, viewModelPageUsuario)

}

func (m *ManipuladorUsuario) CriarUsuarioPost(w http.ResponseWriter, r *http.Request) {

	var usuario = entidades.Usuario{
		Nome:  r.PostFormValue("nome"),
		Email: r.PostFormValue("email"),
		Senha: r.PostFormValue("senha"),
	}

	err := m.servicoUsuario.CriaUsuario(&usuario)

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

	err := m.servicoUsuario.CriarCliente(&cliente)
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

	err := m.servicoUsuario.CriarColaborador(&colaborador)
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

		apresentacao.ExibirErro(w, fmt.Sprintf("Erro manipulador.usuario.editarUsuarioPost:%v", err))
	}

	var usuario = entidades.Usuario{
		UUID:  UUID,
		Nome:  r.PostFormValue("nome"),
		Email: r.PostFormValue("email"),
		Senha: r.PostFormValue("senha"),
	}

	err = m.servicoUsuario.EditarUsuario(usuario)

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

	err = m.servicoUsuario.DeletarUsuario(usuarioUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Usuario:%v", err))
		return
	}

	http.Redirect(w, r, "/usuarios", http.StatusSeeOther)
}
