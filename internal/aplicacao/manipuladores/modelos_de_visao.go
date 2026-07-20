package manipuladores

import (
	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
)

type ViewModelBase struct {
	UsuarioUUID string `json:"usuario_uuid,omitempty"`
}

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

// response DTO
type ViewModelTarefa struct {
	UUID            string      `json:"uuid,omitempty"`
	ProcessoUUID    string      `json:"processo_uuid,omitempty"`
	ResponsavelUUID string      `json:"responsavel_uuid,omitempty"`
	Tarefa          interface{} `json:"tarefa,omitempty"`
	Usuarios        interface{} `json:"usuarios,omitempty"`
}
