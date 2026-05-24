package usuarios

import (
	"net/http"
)

//func (m *ManipuladorUsuario) Rotas(handlerTarefas http.HandlerFunc) *http.ServeMux

func (m *ManipuladorUsuario) DefinirRotasUsuarios() *http.ServeMux {

	// ROTEADOR:ROTEADOR: USUARIOS
	usuariosMux := http.NewServeMux()
	usuariosMux.HandleFunc("GET /{$}", m.PageListarUsuarios)
	usuariosMux.HandleFunc("GET /criar", m.PageCriarUsuario)
	usuariosMux.HandleFunc("GET /{UUID}", m.PageVerUsuario)
	usuariosMux.HandleFunc("GET /{UUID}/editar", m.PageEditarUsuario)
	usuariosMux.HandleFunc("POST /criar", m.CriarUsuarioPost)
	usuariosMux.HandleFunc("POST /{UUID}/editar", m.EditarUsuarioPost)
	usuariosMux.HandleFunc("POST /{UUID}/deletar", m.DeletarUsuarioPost)
	//mux.HandleFunc("POST /{UUID}/tarefas", handlerTarefas)

	return usuariosMux
}

func (m *ManipuladorUsuario) DefinirRotasClientes() *http.ServeMux {
	// ROTEADOR: CLIENTES
	clientesMux := http.NewServeMux()
	clientesMux.HandleFunc("GET /{$}", m.PageListarClientes)
	clientesMux.HandleFunc("GET /criar", m.PageCriarCliente)
	//clientesMux.HandleFunc("GET /{UUID}", m.PageVerCliente)
	//clientesMux.HandleFunc("GET /{UUID}/editar", m.PageEditarCliente)
	clientesMux.HandleFunc("POST /criar", m.CriarClientePost)
	//clientesMux.HandleFunc("POST /{UUID}/editar", m.EditarClientePost)
	//clientesMux.HandleFunc("POST /{UUID}/deletar", m.DeletarClientePost)

	return clientesMux
}

func (m *ManipuladorUsuario) DefinirRotasColaboradores() *http.ServeMux {
	// ROTEADOR: COLABORADORES
	colaboradoresMux := http.NewServeMux()
	colaboradoresMux.HandleFunc("GET /{$}", m.PageListarColaboradores)
	colaboradoresMux.HandleFunc("GET /criar", m.PageCriarColaborador)
	colaboradoresMux.HandleFunc("POST /criar", m.CriarColaboradorPost)

	return colaboradoresMux
}
