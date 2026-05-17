package roteamento

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao/casosdeuso"
	"github.com/gleberphant/ProcessoMan/internal/dominios/paginasestaticas"
	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"
	casosdeuso1 "github.com/gleberphant/ProcessoMan/internal/dominios/usuarios/casosdeuso"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/bancodedados"
	"github.com/gleberphant/ProcessoMan/internal/intermediarios"
)

type Roteador struct {
	LoginManipulador   *autenticacao.ManipuladorLogin
	ManipuladorUsuario *usuarios.ManipuladorUsuario
}

// configurar as rotas e devolver MUX configurado
func (s *Roteador) ConfigurarRotas() http.Handler {

	// configurar mux
	mux := http.NewServeMux()

	// rotas
	// AUTENTICACAO
	// -- login formulario - formulario para enviar usuario e senha
	mux.HandleFunc("GET /login", s.LoginManipulador.PageLogin)
	// -- login acao - recebe usuario e senha para devolver token
	mux.HandleFunc("POST /login", s.LoginManipulador.LoginPost)

	mux.HandleFunc("GET /", paginasestaticas.Index)

	// rotas para os manipuladores usuarios
	mux.HandleFunc("GET /usuario/listar", s.ManipuladorUsuario.PageListar)
	mux.HandleFunc("GET /usuario/criar", s.ManipuladorUsuario.PageCriar)
	mux.HandleFunc("GET /usuario/editar", s.ManipuladorUsuario.PageCriar)

	mux.HandleFunc("POST /usuario/criar", s.ManipuladorUsuario.CriarUsuarioPost)
	mux.HandleFunc("POST /usuario/deletar", s.ManipuladorUsuario.DeletarUsuarioPost)
	mux.HandleFunc("POST /usuario/editar", s.ManipuladorUsuario.EditarUsuarioPost)

	// // rotas para os manipuladores processos

	mux.HandleFunc("GET /processo/listar", s.ManipuladorUsuario.PageListar)
	mux.HandleFunc("GET /processo/criar", s.ManipuladorUsuario.PageCriar)
	mux.HandleFunc("GET /processo/editar", s.ManipuladorUsuario.PageCriar)

	mux.HandleFunc("POST /processo/criar", s.ManipuladorUsuario.CriarUsuarioPost)
	mux.HandleFunc("POST /processo/deletar", s.ManipuladorUsuario.DeletarUsuarioPost)
	mux.HandleFunc("POST /processo/editar", s.ManipuladorUsuario.EditarUsuarioPost)

	return intermediarios.AuthMiddleware(intermediarios.LogMiddleware(mux), s.LoginManipulador.CDUAutenticacao) //-- para auth global

}

func (s *Roteador) InjetarDependencias() error {

	// conexao com o banco de dados
	db, err := bancodedados.ConectarSQLITE()

	if err != nil {
		return err
	}

	// cria os repositorios
	tokensRepo := autenticacao.NovoRepositorioToken(db)
	usuariosRepo := usuarios.NovoRepositorioUsuario(db)

	// injeta repositorios nos casos de uso
	cduAutenticacao := casosdeuso.NovoCasoDeUsoAutenticacao(tokensRepo, usuariosRepo)
	cduUsuario := casosdeuso1.NovoCasoDeUsoUsuario(usuariosRepo)

	// injeta casos de uso nos manipuladores
	s.LoginManipulador = autenticacao.NovoManipuladorLogin(cduAutenticacao, cduUsuario)
	s.ManipuladorUsuario = usuarios.NovoManipuladorUsuario(cduUsuario)

	return nil

}
