package usuarios

type ViewModelUsuario struct {
	UUID     string      `json:"uuid,omitempty"`
	Usuarios interface{} `json:"usuarios,omitempty"`
	Tarefas  interface{} `json:"tarefas,omitempty"`
	Anexos   interface{} `json:"anexos,omitempty"`
}
