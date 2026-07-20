package manipuladores

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/gleberphant/ProcessoMan/internal/dominio/servicos"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

// ManipuladorTarefa gerencia as requisições HTTP relacionadas ao domínio de Tarefas,
// servindo como interface entre a camada de apresentação e os casos de uso.

type ManipuladorTarefa struct {
	servicoProcesso *servicos.ServicoProcesso
	servicoTarefa   *servicos.ServicoTarefa
	servicoUsuario  *servicos.ServicoUsuario
}

// NovoManipuladorTarefa cria e retorna uma nova instância de ManipuladorTarefa.
func NovoManipuladorTarefa(servicoTarefa *servicos.ServicoTarefa,
	servicoUsuario *servicos.ServicoUsuario,
	servicoProcesso *servicos.ServicoProcesso,
) *ManipuladorTarefa {
	return &ManipuladorTarefa{
		servicoTarefa:   servicoTarefa,
		servicoUsuario:  servicoUsuario,
		servicoProcesso: servicoProcesso,
	}
}

func (m *ManipuladorTarefa) Fechar() {
	m.servicoTarefa.Fechar()
	m.servicoUsuario.Fechar()
}

// PageCriar renderiza o formulário para criação de um novo Tarefa.
func (m *ManipuladorTarefa) PageCriarTarefa(w http.ResponseWriter, r *http.Request) {

	var strProcessoUUID string
	var strResponsavelUUID string
	var processoUUID uuid.UUID = uuid.Nil
	var responsavelUUID uuid.UUID = uuid.Nil
	var err error

	strProcessoUUID = r.PostFormValue("ProcessoUUID")
	if strProcessoUUID != "" {

		processoUUID, err = uuid.Parse(strProcessoUUID)

		if err != nil {
			apresentacao.ExibirErro(w, fmt.Sprintf("UUID do processo inválido: %v", err))
			return
		}

		err = m.servicoTarefa.ValidarProcesso(processoUUID)

		if err != nil {
			apresentacao.ExibirErro(w, fmt.Sprintf("Processo invalido:%v", err))
			return
		}
	}

	strResponsavelUUID = r.PostFormValue("ResponsavelUUID")

	if strResponsavelUUID != "" {
		responsavelUUID, err = uuid.Parse(strResponsavelUUID)

		if err != nil {
			apresentacao.ExibirErro(w, fmt.Sprintf("UUID do responsavel inválido: %v", err))
			return
		}

	}

	listaUsuario, err := m.servicoUsuario.ListarUsuarios()

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao carregar usuarios: %v", err))
		return
	}

	listaProcesso, err := m.servicoProcesso.ListarProcessos()

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao carregar processo: %v", err))
		return
	}

	viewModel := struct {
		Tarefa    entidades.Tarefa
		Usuarios  []entidades.Usuario
		Processos []entidades.Processo
	}{
		Tarefa: entidades.Tarefa{
			UUID:            uuid.Nil,
			ProcessoUUID:    processoUUID,
			ResponsavelUUID: responsavelUUID,
		},
		Usuarios:  listaUsuario,
		Processos: listaProcesso,
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-criar-tarefa.html", w, r, viewModel)
}

// PageListar renderiza a página contendo a listagem de todos os Tarefas.
func (m *ManipuladorTarefa) PageListarTarefas(w http.ResponseWriter, r *http.Request) {

	lista, err := m.servicoTarefa.ListarTarefas()

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Listar Tarefa:%v", err))
		return
	}

	viewModel := struct {
		ProcessoUUID string
		Tarefas      []entidades.Tarefa
	}{
		ProcessoUUID: "",
		Tarefas:      lista,
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-listar-tarefas.html", w, r, viewModel)

}

func (m *ManipuladorTarefa) PageVerTarefa(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	tarefaUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID da tarefa inválido: %v", err))
		return
	}

	tarefa, err := m.servicoTarefa.BuscarTarefaPorUUID(tarefaUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("erro PageVerTarefa:%v", err))
		return
	}

	viewModel := struct {
		UUID   string
		Tarefa *entidades.Tarefa
	}{
		UUID:   strUUID,
		Tarefa: tarefa,
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-ver-tarefa.html", w, r, viewModel)
}

// PageEditar carrega os dados de um Tarefa existente e renderiza o mesmo formulário.
func (m *ManipuladorTarefa) PageEditarTarefa(w http.ResponseWriter, r *http.Request) {
	strUUID := r.PathValue("UUID")

	tarefaUUID, err := uuid.Parse(strUUID)
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("UUID da tarefa inválido: %v", err))
		return
	}

	tarefa, err := m.servicoTarefa.BuscarTarefaPorUUID(tarefaUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Editar Tarefa:%v", err))
		return
	}

	listaUsuario, err := m.servicoUsuario.ListarUsuarios()
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ao carregar usuarios: %v", err))
		return
	}

	viewModel := struct {
		UUID     string
		Tarefa   *entidades.Tarefa
		Usuarios []entidades.Usuario
	}{
		UUID:     strUUID,
		Tarefa:   tarefa,
		Usuarios: listaUsuario,
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-criar-tarefa.html", w, r, viewModel)
}

// --------
func (m *ManipuladorTarefa) CriarTarefaPost(w http.ResponseWriter, r *http.Request) {

	processoUUID, err := uuid.Parse(r.PostFormValue("ProcessoUUID"))

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ProcessoUUID: %v", err))
		return
	}

	responsavelUUID, err := uuid.Parse(r.PostFormValue("ResponsavelUUID"))

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro ResponsavelUUID:%v", err))
		return
	}

	tarefa := entidades.Tarefa{
		ProcessoUUID:    processoUUID,
		ResponsavelUUID: responsavelUUID,
		Nome:            r.PostFormValue("Nome"),
		Comentarios:     r.PostFormValue("Comentarios"),
	}

	err = m.servicoTarefa.CriarTarefa(tarefa)

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

	if err := m.servicoTarefa.EditarTarefa(tarefa); err != nil {
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

	err = m.servicoTarefa.DeletarTarefa(tarefaUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Deletar Tarefa: %v", err))
		return
	}

	apresentacao.RedirecionarPaginaAnterior(w, r)
}
