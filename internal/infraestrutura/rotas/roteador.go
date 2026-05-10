package rotas

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/controladores"
	"github.com/gleberphant/ProcessoMan/internal/servicos"
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

	mux.Handle("GET /", servicos.AuthMiddleware(http.HandlerFunc(controladores.Index)))

	// retornar mux
	return servicos.LogMiddleware(mux)

}
