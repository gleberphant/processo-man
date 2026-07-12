package tarefas

import "net/http"

func (m *ManipuladorTarefa) InjetarRotas(mux *http.ServeMux) {

	// SUB-ROTEADOR: TAREFAS
	mux.HandleFunc("GET /tarefas/{$}", m.PageListarTarefas)
	mux.HandleFunc("GET /tarefas/criar", m.PageCriarTarefa)
	mux.HandleFunc("GET /tarefas/{UUID}", m.PageVerTarefa)
	mux.HandleFunc("GET /tarefas/{UUID}/editar", m.PageEditarTarefa)

	mux.HandleFunc("POST /tarefas/criar", m.CriarTarefaPost)
	mux.HandleFunc("POST /tarefas/{UUID}/editar", m.EditarTarefaPost)
	mux.HandleFunc("POST /tarefas/{UUID}/deletar", m.DeletarTarefaPost)
	mux.HandleFunc("POST /tarefas/{UUID}/concluir", m.DeletarTarefaPost)

	//mux.HandleFunc("GET /processos/{processo_uuid}/tarefas", m.PageListarTarefasPorProcesso)
	mux.HandleFunc("GET /processos/{colaborador_uuid}/tarefas", m.PageListarTarefasPorResponsavel)
}
