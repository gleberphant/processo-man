package processos

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
)

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
