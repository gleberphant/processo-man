package roteamento

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
)

// configurar as rotas e devolver MUX configurado
func (s *Roteador) InjetarRotas() {

	// configurar mux
	mux := http.NewServeMux()

	// PAGINAS ESTÁTICAS
	mux.HandleFunc("/{$}", s.Index)
	mux.HandleFunc("/", s.Pagina404)

	s.ManipuladorAutenticacao.DefinirRotasAutenticacao(mux)
	s.ManipuladorProcesso.InjetarRotasProcessos(mux)
	s.ManipuladorUsuario.InjetarRotasUsuarios(mux)
	s.ManipuladorUsuario.InjetarRotasClientes(mux)
	s.ManipuladorUsuario.InjetarRotasColaboradores(mux)
	s.ManipuladorTarefa.InjetarRotasTarefas(mux)

	// INJETA INTERMEDIÁRIOS - Middlewares
	roteador := autenticacao.LogIntermediario(autenticacao.AutenticadorIntermediario(mux, s.ManipuladorAutenticacao.CDUAutenticacao))

	s.Handler = &roteador

}
