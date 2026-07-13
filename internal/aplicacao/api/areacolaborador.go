package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

type ICDUUsuario interface {
	ListarClientes() ([]entidades.Cliente, error)
}

type ICDUProcesso interface {
	ListarProcessosPorCliente(uuid.UUID) ([]entidades.Processo, error)
	BuscarProcessoPorUUID(uuid.UUID) (*entidades.Processo, error)
}

// servindo como interface entre a camada de apresentação e os casos de uso.
type ManipuladorProcessoApi struct {
	cduProcesso ICDUProcesso
	cduUsuario  ICDUUsuario
}

func (m *ManipuladorProcessoApi) APIVisualizarProcesso(w http.ResponseWriter, r *http.Request) {
	uuidStr := r.URL.Query().Get("uuid")

	processoUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		erroMsg := fmt.Sprintf("UUID do processo inválido: %v", err)
		log.Println(erroMsg)
		http.Error(w, erroMsg, http.StatusBadRequest)
		return
	}

	processo, err := m.cduProcesso.BuscarProcessoPorUUID(processoUUID)

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
			Processo entidades.Processo
			Anexos   []string
			Tarefas  []entidades.Tarefa
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
