package usuarios

import (
	"net/http"
)

//func (m *ManipuladorUsuario) Rotas(handlerTarefas http.HandlerFunc) *http.ServeMux

func (m *ManipuladorUsuario) InjetarRotasUsuarios(mux *http.ServeMux) {
	// ROTAS DE USUÁRIOS
	mux.HandleFunc("GET /usuarios/{$}", m.PageListarUsuarios)
	mux.HandleFunc("GET /usuarios/criar", m.PageCriarUsuario)
	mux.HandleFunc("GET /usuarios/{UUID}", m.PageVerUsuario)
	mux.HandleFunc("GET /usuarios/{UUID}/editar", m.PageEditarUsuario)
	mux.HandleFunc("POST /usuarios/criar", m.CriarUsuarioPost)
	mux.HandleFunc("POST /usuarios/{UUID}/editar", m.EditarUsuarioPost)
	mux.HandleFunc("POST /usuarios/{UUID}/deletar", m.DeletarUsuarioPost)
}

func (m *ManipuladorUsuario) InjetarRotasClientes(mux *http.ServeMux) {
	// ROTAS DE CLIENTES
	mux.HandleFunc("GET /usuarios/clientes/{$}", m.PageListarClientes)
	mux.HandleFunc("GET /usuarios/clientes/criar", m.PageCriarCliente)
	mux.HandleFunc("POST /usuarios/clientes/criar", m.CriarClientePost)
}

func (m *ManipuladorUsuario) InjetarRotasColaboradores(mux *http.ServeMux) {
	// rotas da area do colaborador -- temporario depois mudar para o local correto
	mux.HandleFunc("GET /colaboradores/{$}", m.PageListarColaboradores)
	mux.HandleFunc("GET /colaboradores/criar", m.PageCriarColaborador)
	mux.HandleFunc("POST /colaboradores/criar", m.CriarColaboradorPost)

	// ROTAS DE GESTÃO DOS COLABORADORES
	mux.HandleFunc("GET /usuarios/colaboradores/{$}", m.PageListarColaboradores)
	mux.HandleFunc("GET /usuarios/colaboradores/criar", m.PageCriarColaborador)
	mux.HandleFunc("POST /usuarios/colaboradores/criar", m.CriarColaboradorPost)
}
