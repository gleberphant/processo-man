package manipuladores

import (
	"net/http"
)

//func (m *ManipuladorUsuario) Rotas(handlerTarefas http.HandlerFunc) *http.ServeMux

func (m *ManipuladorUsuario) InjetarRotas(mux *http.ServeMux) {

	// ROTAS DE USUÁRIOS
	mux.HandleFunc("GET /usuarios/{$}", m.PageListarUsuarios)
	mux.HandleFunc("GET /usuarios/criar", m.PageCriarUsuario)
	mux.HandleFunc("GET /usuarios/{UUID}", m.PageVerUsuario)
	mux.HandleFunc("GET /usuarios/{UUID}/editar", m.PageEditarUsuario)
	mux.HandleFunc("POST /usuarios/criar", m.CriarUsuarioPost)
	mux.HandleFunc("POST /usuarios/{UUID}/editar", m.EditarUsuarioPost)
	mux.HandleFunc("POST /usuarios/{UUID}/deletar", m.DeletarUsuarioPost)

	// ROTAS DE CLIENTES
	mux.HandleFunc("GET /usuarios/clientes/{$}", m.PageListarClientes)
	mux.HandleFunc("GET /usuarios/clientes/criar", m.PageCriarCliente)
	mux.HandleFunc("POST /usuarios/clientes/criar", m.CriarClientePost)
	mux.HandleFunc("GET /clientes/{$}", m.PageListarClientes)

	// rotas da area do colaborador -- temporario depois mudar para o local correto
	mux.HandleFunc("GET /colaboradores/{$}", m.PageListarColaboradores)
	mux.HandleFunc("GET /colaboradores/criar", m.PageCriarColaborador)
	mux.HandleFunc("POST /colaboradores/criar", m.CriarColaboradorPost)

	// ROTAS DE GESTÃO DOS COLABORADORES
	mux.HandleFunc("GET /usuarios/colaboradores/{$}", m.PageListarColaboradores)
	mux.HandleFunc("GET /usuarios/colaboradores/criar", m.PageCriarColaborador)
	mux.HandleFunc("POST /usuarios/colaboradores/criar", m.CriarColaboradorPost)
}
