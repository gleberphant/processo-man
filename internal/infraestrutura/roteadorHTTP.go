package infraestrutura

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/casosdeuso/intermediarios"
	"github.com/gleberphant/ProcessoMan/internal/manipuladores"
)

type Roteador struct {
}

// configurar as rotas e devolver MUX configurado
func (s *Roteador) ConfigurarRotas() http.Handler {

	// configurar mux
	mux := http.NewServeMux()

	// rotas
	// AUTENTICACAO
	// -- login formulario - formulario para enviar usuario e senha
	mux.HandleFunc("GET /login", manipuladores.LoginGet)
	// -- login acao - recebe usuario e senha para devolver token
	mux.HandleFunc("POST /login", manipuladores.LoginPost)

	// ROTAS PROTEGIDAS  APOS AUTENTICADO
	// PASSA PELO MIDDLEWARE AUTENTICADOR
	// -- home -

	// mux.Handle("GET /", intermediarios.AuthMiddleware(http.HandlerFunc(controladores.Index)))      // para auth local
	// mux.Handle("GET /sss", intermediarios.AuthMiddleware(http.HandlerFunc(controladores.Page1)))   // para auth local
	// mux.Handle("GET /page2", intermediarios.AuthMiddleware(http.HandlerFunc(controladores.Page2))) // para auth local

	// para auth global
	mux.HandleFunc("GET /", manipuladores.Index)
	mux.HandleFunc("GET /page1", manipuladores.Page1)
	mux.HandleFunc("GET /page2", manipuladores.Page2)

	// retornar mux
	//	return intermediarios.LogMiddleware(mux) // para auth local
	return intermediarios.AuthMiddleware(intermediarios.LogMiddleware(mux)) //-- para auth global

}
