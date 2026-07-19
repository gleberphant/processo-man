package manipuladores

import "net/http"

func (m *ManipuladorAutenticacao) InjetarRotas(mux *http.ServeMux) {
	// AUTENTICACAO
	mux.HandleFunc("GET /login", m.PageLogin)
	mux.HandleFunc("POST /login", m.LoginPost)
}

func (m *ManipuladorUsuario) InjetarRotas(mux *http.ServeMux) {
	// ROTAS DE USUÁRIOS
	mux.HandleFunc("GET /usuarios/{$}", m.PageListarUsuario)
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
	//mux.HandleFunc("GET /clientes/{$}", m.PageListarClientes)

	// ROTAS DE GESTÃO DOS COLABORADORES
	mux.HandleFunc("GET /usuarios/colaboradores/{$}", m.PageListarColaboradores)
	mux.HandleFunc("GET /usuarios/colaboradores/criar", m.PageCriarColaborador)
	mux.HandleFunc("POST /usuarios/colaboradores/criar", m.CriarColaboradorPost)
}

func (m *ManipuladorProcesso) InjetarRotas(mux *http.ServeMux) {
	// ROTAS DE PROCESSOS
	mux.HandleFunc("GET /processos/{$}", m.PageListar)
	mux.HandleFunc("GET /processos/criar", m.PageCriar)
	mux.HandleFunc("GET /processos/{UUID}", m.PageVerProcesso)
	mux.HandleFunc("GET /processos/{UUID}/editar", m.PageEditar)

	mux.HandleFunc("POST /processos/criar", m.CriarProcessoPost)
	mux.HandleFunc("POST /processos/{UUID}/editar", m.EditarProcessoPost)
	mux.HandleFunc("POST /processos/{UUID}/deletar", m.DeletarProcessoPost)
	mux.HandleFunc("POST /processos/{UUID}/arquivar", m.DeletarProcessoPost)

}

func (m *ManipuladorTarefa) InjetarRotas(mux *http.ServeMux) {
	// ROTEADOR: TAREFAS
	mux.HandleFunc("GET /tarefas/{$}", m.PageListarTarefas)
	mux.HandleFunc("GET /tarefas/criar", m.PageCriarTarefa)
	mux.HandleFunc("GET /tarefas/{UUID}", m.PageVerTarefa)
	mux.HandleFunc("GET /tarefas/{UUID}/editar", m.PageEditarTarefa)

	mux.HandleFunc("POST /tarefas/criar", m.CriarTarefaPost)
	mux.HandleFunc("POST /tarefas/{UUID}/editar", m.EditarTarefaPost)
	mux.HandleFunc("POST /tarefas/{UUID}/deletar", m.DeletarTarefaPost)
	mux.HandleFunc("POST /tarefas/{UUID}/concluir", m.DeletarTarefaPost)

}
