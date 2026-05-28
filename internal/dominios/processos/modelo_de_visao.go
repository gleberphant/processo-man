package processos

type ViewModelProcesso struct {
	UUID      string
	Processos interface{}
	Anexos    []string
	Usuarios  interface{}
}

// DTO PROCESSO
type ProcessoResponse struct {
}
