package roteamento

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominios/paginasestaticas"
	"github.com/gleberphant/ProcessoMan/internal/intermediarios"
)

// configurar as rotas e devolver MUX configurado
func (s *Roteador) ConfigurarRotas() http.Handler {

	// configurar mux
	mux := http.NewServeMux()

	// rotas
	// PAGINAS ESTÁTICAS
	mux.HandleFunc("GET /", paginasestaticas.Index)
	mux.HandleFunc("GET /404", paginasestaticas.Index)

	// AUTENTICACAO
	mux.HandleFunc("GET /login", s.LoginManipulador.PageLogin)
	mux.HandleFunc("POST /login", s.LoginManipulador.LoginPost)

	// ROTAS USUARIOS
	//// PAGINAS GET
	mux.HandleFunc("GET /usuarios", s.ManipuladorUsuario.PageListarUsuarios)
	mux.HandleFunc("GET /usuarios/criar", s.ManipuladorUsuario.PageCriarUsuario)
	mux.HandleFunc("GET /usuarios/{UUID}", s.ManipuladorUsuario.PageVerUsuario)
	mux.HandleFunc("GET /usuarios/{UUID}/editar", s.ManipuladorUsuario.PageEditarUsuario)

	//// ACOES POST
	mux.HandleFunc("POST /usuarios/criar", s.ManipuladorUsuario.CriarUsuarioPost)
	mux.HandleFunc("POST /usuarios/{UUID}/editar", s.ManipuladorUsuario.EditarUsuarioPost)
	mux.HandleFunc("POST /usuarios/{UUID}/deletar", s.ManipuladorUsuario.DeletarUsuarioPost)

	// ROTAS PROCESSOS
	//// PAGINAS GET
	mux.HandleFunc("GET /processos", s.ManipuladorProcesso.PageListar)
	mux.HandleFunc("GET /processos/criar", s.ManipuladorProcesso.PageCriar)
	mux.HandleFunc("GET /processos/{UUID}", s.ManipuladorProcesso.PageVerProcesso)
	mux.HandleFunc("GET /processos/{UUID}/editar", s.ManipuladorProcesso.PageEditar)
	mux.HandleFunc("GET /processos/{UUID}/tarefas", s.ManipuladorTarefa.PageListarTarefasPorProcesso)

	//// AÇÕES POST
	mux.HandleFunc("POST /processos/criar", s.ManipuladorProcesso.CriarProcessoPost)
	mux.HandleFunc("POST /processos/{UUID}/editar", s.ManipuladorProcesso.EditarProcessoPost)
	mux.HandleFunc("POST /processos/{UUID}/deletar", s.ManipuladorProcesso.DeletarProcessoPost)
	mux.HandleFunc("POST /processos/{UUID}/arquivar", s.ManipuladorProcesso.DeletarProcessoPost)

	// ROTAS TAREFAS
	//// PAGINAS GET
	mux.HandleFunc("GET /tarefas", s.ManipuladorTarefa.PageListarTarefas)
	mux.HandleFunc("GET /tarefas/criar", s.ManipuladorTarefa.PageCriarTarefa)
	mux.HandleFunc("GET /tarefas/{UUID}", s.ManipuladorTarefa.PageVerTarefa)
	mux.HandleFunc("GET /tarefas/{UUID}/editar", s.ManipuladorTarefa.PageEditarTarefa)

	//// AÇÕES POST
	mux.HandleFunc("POST /tarefas/criar", s.ManipuladorTarefa.CriarTarefaPost)
	mux.HandleFunc("POST /tarefas/{UUID}/editar", s.ManipuladorTarefa.EditarTarefaPost)
	mux.HandleFunc("POST /tarefas/{UUID}/deletar", s.ManipuladorTarefa.DeletarTarefaPost)
	mux.HandleFunc("POST /tarefas/{UUID}/concluir", s.ManipuladorTarefa.DeletarTarefaPost)

	//// RETORNOS DE API
	//mux.HandleFunc("GET /api/processo/visualizar", s.ManipuladorProcesso.APIVisualizarProcesso)

	// INJETA INTERMEDIÁRIOS - Middlewares
	roteador := intermediarios.AutenticadorIntermediario(intermediarios.LogIntermediario(mux), s.LoginManipulador.CDUAutenticacao)
	return roteador

}
