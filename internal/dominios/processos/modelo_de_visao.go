package processos

import "github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"

type ViewModelProcesso struct {
	apresentacao.BaseViewModel
	UUID      string
	Processos interface{}
	Anexos    []string
	Usuarios  interface{}
}

// DTO PROCESSO
type ProcessoResponse struct {
}
