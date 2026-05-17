package processos

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominios/processos/casosdeuso"
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

// ManipuladorProcesso gerencia as requisições HTTP relacionadas ao domínio de processos,
// servindo como interface entre a camada de apresentação e os casos de uso.
type ManipuladorProcesso struct {
	cduProcesso *casosdeuso.CDUProcesso
}

// NovoManipuladorProcesso cria e retorna uma nova instância de ManipuladorProcesso.
func NovoManipuladorProcesso(CasosDeUsoProcesso *casosdeuso.CDUProcesso) *ManipuladorProcesso {
	return &ManipuladorProcesso{
		cduProcesso: CasosDeUsoProcesso,
	}
}

// PageListar renderiza a página contendo a listagem de todos os processos.
func (m *ManipuladorProcesso) PageListar(w http.ResponseWriter, r *http.Request) {

	lista, err := m.cduProcesso.ListarProcessos()

	if err != nil {
		erro := fmt.Sprintf("Erro :%v", err)
		log.Println(erro)
		//substituir por redirecionamento para o index com uma mensagem
		http.Error(w, erro, http.StatusInternalServerError)
	}

	apresentacao.ExibirPaginaHTML("processo/page-listar-processos.html", w, lista)

}

// PageCriar renderiza o formulário para criação de um novo processo.
func (m *ManipuladorProcesso) PageCriar(w http.ResponseWriter, r *http.Request) {
	// Passa uma entidade vazia para o template para manter a consistência com a edição
	apresentacao.ExibirPaginaHTML("processo/page-criar-processo.html", w, entidades.Processo{})
}

// PageEditar carrega os dados de um processo existente e renderiza o mesmo formulário.
func (m *ManipuladorProcesso) PageEditar(w http.ResponseWriter, r *http.Request) {
	// Assumindo passagem do UUID via Query String (?uuid=...). Ajuste caso use chi.URLParam(r, "uuid")
	uuidStr := r.URL.Query().Get("uuid")

	// Nota: Utilizando o método existente. Recomenda-se renomear BuscarUsuarioPorUUID para BuscarProcessoPorUUID no futuro.
	processo, err := m.cduProcesso.BuscarProcessoPorUUID(uuidStr)
	if err != nil {
		erroMsg := fmt.Sprintf("Processo não encontrado: %v", err)
		log.Println(erroMsg)
		http.Error(w, erroMsg, http.StatusNotFound)
		return
	}

	// Reutiliza o mesmo template, injetando os dados do processo
	apresentacao.ExibirPaginaHTML("processo/page-criar-processo.html", w, processo)
}

// CriarProcessoPost processa a submissão do formulário para persistir um novo processo.
func (m *ManipuladorProcesso) CriarProcessoPost(w http.ResponseWriter, r *http.Request) {

	UUID, err := uuid.Parse(r.PostFormValue("uuid"))
	var Processo = entidades.Processo{
		UUID: UUID,
		Nome: r.PostFormValue("nome"),
		ListaTarefas: []entidades.Tarefa{{
			Nome: r.PostFormValue("tarefa"),
		},
		},
	}

	err = m.cduProcesso.CriarProcesso(Processo)

	if err != nil {
		erroMsg := fmt.Sprintf("Erro na criação do Processo:%v", err)
		log.Println(erroMsg)
		//substituir por redirecionamento para o index com uma mensagem
		http.Error(w, erroMsg, http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/processo/listar", http.StatusSeeOther)

}

// EditarProcessoPost processa a atualização de um processo existente.
func (m *ManipuladorProcesso) EditarProcessoPost(w http.ResponseWriter, r *http.Request) {

	UUID, err := uuid.Parse(r.PostFormValue("uuid"))
	var Processo = entidades.Processo{
		UUID: UUID,
		Nome: r.PostFormValue("nome"),
		ListaTarefas: []entidades.Tarefa{{
			Nome: r.PostFormValue("tarefa"),
		},
		},
	}

	err = m.cduProcesso.AtualizarProcesso(Processo)

	if err != nil {
		erroMsg := fmt.Sprintf("Erro na edição do Processo:%v", err)
		log.Println(erroMsg)
		//substituir por redirecionamento para o index com uma mensagem
		http.Error(w, erroMsg, http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/processo/listar", http.StatusSeeOther)
}

// DeletarProcessoPost remove um processo com base no identificador enviado via formulário.
func (m *ManipuladorProcesso) DeletarProcessoPost(w http.ResponseWriter, r *http.Request) {

	var UUID = r.PostFormValue("uuid")

	err := m.cduProcesso.DeletarProcesso(UUID)

	if err != nil {
		erroMsg := fmt.Sprintf("Erro ao deletar Processo:%v", err)
		log.Println(erroMsg)
		//substituir por redirecionamento para o index com uma mensagem
		http.Error(w, erroMsg, http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/processo/listar", http.StatusSeeOther)
}
