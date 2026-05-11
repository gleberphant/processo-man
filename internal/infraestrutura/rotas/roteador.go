package rotas

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/controladores"
	servicos "github.com/gleberphant/ProcessoMan/internal/servicos/intermediarios"
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
	mux.HandleFunc("GET /login", controladores.FormularioLogin)
	// -- login acao - recebe usuario e senha para devolver token
	mux.HandleFunc("POST /login", controladores.Logar)

	// ROTAS PROTEGIDAS  APOS AUTENTICADO
	// PASSA PELO MIDDLEWARE AUTENTICADOR
	// -- home -

	mux.Handle("GET /", servicos.AuthMiddleware(http.HandlerFunc(controladores.Index))) // para auth local
	//mux.HandleFunc("GET /", controladores.Index) -- para auth global

	// retornar mux
	return servicos.LogMiddleware(mux) // para auth local
	// return servicos.AuthMiddleware(servicos.LogMiddleware(mux)) -- para auth global

}
