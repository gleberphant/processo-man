package tarefas

import "net/http"

func (m *ManipuladorTarefa) DefinirRotasTarefas() *http.ServeMux {

	// SUB-ROTEADOR: TAREFAS
	tarefasMux := http.NewServeMux()
	tarefasMux.HandleFunc("GET /{$}", m.PageListarTarefas)
	tarefasMux.HandleFunc("GET /criar", m.PageCriarTarefa)
	tarefasMux.HandleFunc("GET /{UUID}", m.PageVerTarefa)
	tarefasMux.HandleFunc("GET /{UUID}/editar", m.PageEditarTarefa)
	tarefasMux.HandleFunc("POST /criar", m.CriarTarefaPost)
	tarefasMux.HandleFunc("POST /{UUID}/editar", m.EditarTarefaPost)
	tarefasMux.HandleFunc("POST /{UUID}/deletar", m.DeletarTarefaPost)
	tarefasMux.HandleFunc("POST /{UUID}/concluir", m.DeletarTarefaPost)

	return tarefasMux
}
