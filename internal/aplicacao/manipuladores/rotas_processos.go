package manipuladores

import "net/http"

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
