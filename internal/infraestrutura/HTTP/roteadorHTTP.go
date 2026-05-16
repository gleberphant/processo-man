package HTTP

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/adaptadores/intermediarios"
	"github.com/gleberphant/ProcessoMan/internal/adaptadores/manipuladores"
	"github.com/gleberphant/ProcessoMan/internal/adaptadores/repositorios"
	"github.com/gleberphant/ProcessoMan/internal/casosdeuso/autenticacao"
)

type Roteador struct {
	LoginManipulador *manipuladores.ManipuladorLogin
}

// configurar as rotas e devolver MUX configurado
func (s *Roteador) ConfigurarRotas() http.Handler {

	// configurar mux
	mux := http.NewServeMux()

	// rotas
	// AUTENTICACAO
	// -- login formulario - formulario para enviar usuario e senha
	mux.HandleFunc("GET /login", s.LoginManipulador.LoginGet)
	// -- login acao - recebe usuario e senha para devolver token
	mux.HandleFunc("POST /login", s.LoginManipulador.LoginPost)

	mux.HandleFunc("GET /", manipuladores.Index)
	mux.HandleFunc("GET /page1", manipuladores.Page1)
	mux.HandleFunc("GET /page2", manipuladores.Page2)

	return intermediarios.AuthMiddleware(intermediarios.LogMiddleware(mux), s.LoginManipulador.CDUatenticacao) //-- para auth global

}

func (s *Roteador) InjecaoDependencias() {

	// cria os repositorios
	tokensRepo, err := repositorios.NovoTokenRepo()

	if err != nil {
		return
	}

	usuariosRepo, err := repositorios.NovoUsuarioRepo()

	if err != nil {
		return
	}

	// injeta repositorios nos casos de uso
	cduAutenticacao := autenticacao.NovoAutenticacaoCDU(tokensRepo, usuariosRepo)

	// injeta casos de uso nos manipuladores
	s.LoginManipulador = manipuladores.NovoManipuladorLogin(cduAutenticacao)

}
