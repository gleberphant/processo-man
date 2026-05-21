package processos

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominios/tarefas"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
)

func (m *ManipuladorProcesso) APIVisualizarProcesso(w http.ResponseWriter, r *http.Request) {
	uuidStr := r.URL.Query().Get("uuid")

	processo, err := m.cduProcesso.BuscarProcessoPorUUID(uuidStr)

	if err != nil {
		erroMsg := fmt.Sprintf("Processo não encontrado: %v", err)
		log.Println(erroMsg)
		http.Error(w, erroMsg, http.StatusNotFound)
		return
	}

	responseAPI := struct {
		Msg     string
		Payload any
	}{
		Msg: "OK",
		Payload: struct {
			Processo Processo
			Anexos   []string
			Tarefas  []tarefas.Tarefa
		}{
			Processo: *processo,
			Anexos:   []string{"arquivo1.doc", "arquivo2.doc"},
			Tarefas:  processo.Tarefas,
		},
	}

	err = apresentacao.ExibirJsonApi(w, responseAPI)

	if err != nil {
		erroMsg := fmt.Sprintf("Falha no JASON: %v", err)
		log.Println(erroMsg)
		http.Error(w, erroMsg, http.StatusNotFound)
		return
	}
}
