package processos

import "net/http"

func (m *ManipuladorProcesso) DefinirRotasProcessos() *http.ServeMux {

	// SUB-ROTEADOR: PROCESSOS
	processosMux := http.NewServeMux()
	processosMux.HandleFunc("GET /{$}", m.PageListar)
	processosMux.HandleFunc("GET /criar", m.PageCriar)
	processosMux.HandleFunc("GET /{UUID}", m.PageVerProcesso)
	processosMux.HandleFunc("GET /{UUID}/editar", m.PageEditar)
	//processosMux.HandleFunc("GET /{UUID}/tarefas", s.ManipuladorTarefa.PageListarTarefasPorProcesso)
	processosMux.HandleFunc("POST /criar", m.CriarProcessoPost)
	processosMux.HandleFunc("POST /{UUID}/editar", m.EditarProcessoPost)
	processosMux.HandleFunc("POST /{UUID}/deletar", m.DeletarProcessoPost)
	processosMux.HandleFunc("POST /{UUID}/arquivar", m.DeletarProcessoPost)

	return processosMux
}
