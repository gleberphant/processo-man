package manipuladores

import (
	"time"

	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

type ViewModelPageUsuario struct {
	//apresentacao.BaseViewModel
	UUID     string      `json:"uuid,omitempty"`
	Usuario  interface{} `json:"usuarios,omitempty"`
	Processo interface{} `json:"processos,omitempty"`
	Tarefa   interface{} `json:"tarefas,omitempty"`
	Arquivo  interface{} `json:"anexos,omitempty"`
}

//	type tarefasView struct {
//		UUID            uuid.UUID
//		ProcessoUUID    uuid.UUID
//		ResponsavelUUID uuid.UUID
//		Nome            string
//		Concluida       bool
//		Comentarios     string
//		DataConclusao   time.Time
//		DataCriacao     time.Time
//	}
type ViewModelProcesso struct {
	apresentacao.BaseViewModel
	UUID          string
	Clientes      []entidades.Cliente
	Colaboradores []entidades.Colaborador
	Processo      entidades.Processo
	Processos     []entidades.Processo
	Anexos        []string
	Usuarios      []entidades.Usuario
}

// DTO PROCESSO
type ProcessoResponse struct {
}

// response DTO
type ViewModelTarefa struct {
	UUID            string      `json:"uuid,omitempty"`
	ProcessoUUID    string      `json:"processo_uuid,omitempty"`
	ResponsavelUUID string      `json:"responsavel_uuid,omitempty"`
	Tarefa          interface{} `json:"tarefa,omitempty"`
	Usuarios        interface{} `json:"usuarios,omitempty"`
}

// usuarioView responseDTO é privado ao pacote tarefas e define os campos exibidos na seleção.
type tarefaView struct {
	UUID            uuid.UUID `json:"uuid,omitempty"`
	ProcessoUUID    uuid.UUID `json:"processo_uuid,omitempty"`
	ResponsavelUUID uuid.UUID `json:"responsavel_uuid,omitempty"`
	Nome            string    `json:"nome,omitempty"`
	Concluida       bool      `json:"concluida,omitempty"`
	Comentarios     string    `json:"comentarios,omitempty"`
	DataConclusao   time.Time `json:"data_conclusao,omitempty"`
	DataCriacao     time.Time `json:"data_criacao,omitempty"`
}

// usuarioView responseDTO é privado ao pacote tarefas e define os campos exibidos na seleção.
type usuarioView struct {
	UUID string `json:"uuid,omitempty"`
	Nome string `json:"nome,omitempty"`
}
