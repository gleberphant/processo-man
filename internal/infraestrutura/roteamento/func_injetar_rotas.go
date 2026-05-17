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
	mux.HandleFunc("GET /usuario/listar", s.ManipuladorUsuario.PageListar)
	mux.HandleFunc("GET /usuario/criar", s.ManipuladorUsuario.PageCriar)
	mux.HandleFunc("GET /usuario/editar", s.ManipuladorUsuario.PageCriar)

	//// ACOES POST
	mux.HandleFunc("POST /usuario/criar", s.ManipuladorUsuario.CriarUsuarioPost)
	mux.HandleFunc("POST /usuario/deletar", s.ManipuladorUsuario.DeletarUsuarioPost)
	mux.HandleFunc("POST /usuario/editar", s.ManipuladorUsuario.EditarUsuarioPost)

	// ROTAS PROCESSOS
	//// PAGINAS GET
	mux.HandleFunc("GET /processo/listar", s.ManipuladorProcesso.PageListar)
	mux.HandleFunc("GET /processo/criar", s.ManipuladorProcesso.PageCriar)
	mux.HandleFunc("GET /processo/editar", s.ManipuladorProcesso.PageEditar)

	//// AÇÕES POST
	mux.HandleFunc("POST /processo/criar", s.ManipuladorProcesso.CriarProcessoPost)
	mux.HandleFunc("POST /processo/deletar", s.ManipuladorProcesso.DeletarProcessoPost)
	mux.HandleFunc("POST /processo/editar", s.ManipuladorProcesso.EditarProcessoPost)

	// INJETA INTERMEDIÁRIOS - Middlewares
	roteador := intermediarios.AutenticadorIntermediario(intermediarios.LogIntermediario(mux), s.LoginManipulador.CDUAutenticacao)
	return roteador

}
