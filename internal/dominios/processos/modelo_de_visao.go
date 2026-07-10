package processos

import "github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"

type ViewModelProcesso struct {
	apresentacao.BaseViewModel
	UUID      string
	Processo  interface{}
	Processos interface{}
	Anexos    []string
	Usuarios  interface{}
}

// DTO PROCESSO
type ProcessoResponse struct {
}
