package roteamento

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
)

// configurar as rotas e devolver MUX configurado
func (r *Roteador) InjetarRotas() {

	// configurar mux
	mux := http.NewServeMux()

	// PAGINAS ESTÁTICAS
	mux.HandleFunc("/{$}", r.Index)
	mux.HandleFunc("/", r.Pagina404)

	r.ManipuladorAutenticacao.DefinirRotasAutenticacao(mux)
	r.ManipuladorProcesso.InjetarRotasProcessos(mux)
	r.ManipuladorUsuario.InjetarRotasUsuarios(mux)
	r.ManipuladorUsuario.InjetarRotasClientes(mux)
	r.ManipuladorUsuario.InjetarRotasColaboradores(mux)
	r.ManipuladorTarefa.InjetarRotasTarefas(mux)

	// INJETA INTERMEDIÁRIOS - Middlewares
	roteador := autenticacao.AutenticadorIntermediario(mux, r.ManipuladorAutenticacao.CDUAutenticacao)
	roteador = autenticacao.LogIntermediario(roteador)

	r.Handler = &roteador

}
