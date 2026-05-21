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
		erro := fmt.Sprintf("Erro :%v", err)
		log.Println(erro)
		http.Error(w, erro, http.StatusInternalServerError) //substituir por redirecionamento para o index com uma mensagem
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-listar-tarefas.html", w, lista)

}

// PageEditar carrega os dados de um Tarefa existente e renderiza o mesmo formulário.
func (m *ManipuladorTarefa) PageEditarTarefa(w http.ResponseWriter, r *http.Request) {
	uuidStr := r.URL.Query().Get("uuid")

	Tarefa, err := m.cduTarefa.BuscarTarefaPorUUID(uuidStr)
	if err != nil {
		erroMsg := fmt.Sprintf("Tarefa não encontrado: %v", err)
		log.Println(erroMsg)
		http.Error(w, erroMsg, http.StatusNotFound)
		return
	}

	apresentacao.ExibirPaginaHTML("tarefa/page-criar-Tarefa.html", w, Tarefa)
}

// --------
func (m *ManipuladorTarefa) CriarTarefaPost(w http.ResponseWriter, r *http.Request) {
	processoUUID, err := uuid.Parse(r.PostFormValue("ProcessoUUID"))
	if err != nil {
		erroMsg := fmt.Sprintf("Erro na criação do Tarefa:%v", err)
		log.Println(erroMsg)
		http.Error(w, erroMsg, http.StatusInternalServerError) //substituir por redirecionamento para o index com uma mensagem
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
		erroMsg := fmt.Sprintf("Erro na criação do Tarefa:%v", err)
		log.Println(erroMsg)
		http.Error(w, erroMsg, http.StatusInternalServerError) //substituir por redirecionamento para o index com uma mensagem
		return
	}

	http.Redirect(w, r, "/processo/visualizar?uuid="+tarefa.ProcessoUUID.String(), http.StatusSeeOther)

}

func (m *ManipuladorTarefa) EditarTarefaPost(w http.ResponseWriter, r *http.Request) {

	UUID, err := uuid.Parse(r.PostFormValue("uuid"))

	var Tarefa = Tarefa{
		UUID:            UUID,
		ProcessoUUID:    uuid.MustParse(r.PostFormValue("ProcessoUUID")),
		ResponsavelUUID: uuid.MustParse(r.PostFormValue("ResponsavelUUID")),
		Nome:            r.PostFormValue("Nome"),
		Comentarios:     r.PostFormValue("Comentarios"),
	}

	err = m.cduTarefa.EditarTarefa(Tarefa)

	if err != nil {
		erroMsg := fmt.Sprintf("Erro na edição do Tarefa:%v", err)
		log.Println(erroMsg)

		http.Error(w, erroMsg, http.StatusInternalServerError) //substituir por redirecionamento para o index com uma mensagem
	}

	http.Redirect(w, r, "/tarefa/listar", http.StatusSeeOther)
}

func (m *ManipuladorTarefa) DeletarTarefaPost(w http.ResponseWriter, r *http.Request) {

	var UUID = r.PostFormValue("uuid")

	err := m.cduTarefa.DeletarTarefa(UUID)

	if err != nil {
		erroMsg := fmt.Sprintf("Erro ao deletar Tarefa:%v", err)
		log.Println(erroMsg)

		http.Error(w, erroMsg, http.StatusInternalServerError) //substituir por redirecionamento para o index com uma mensagem
	}

	http.Redirect(w, r, "/Tarefa/listar", http.StatusSeeOther)
}
