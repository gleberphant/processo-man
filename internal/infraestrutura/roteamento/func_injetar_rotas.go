package roteamento

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/intermediarios"
	"github.com/gleberphant/ProcessoMan/internal/manipuladores"
)

// configurar as rotas e devolver MUX configurado
func (r *Roteador) InjetarRotas() {

	// configurar mux
	mux := http.NewServeMux()

	// PAGINAS ESTÁTICAS
	// index
	mux.HandleFunc("/{$}", manipuladores.Index)

	// ferramentas dev do google chrome
	mux.HandleFunc("/.well-known/*", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	// retornar o favicon
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/favicon.ico") // Ajuste o caminho do seu arquivo
	})

	// ROTAS DOS MANIPULADORES
	r.ManipuladorAutenticacao.InjetarRotas(mux)
	r.ManipuladorProcesso.InjetarRotas(mux)
	r.ManipuladorUsuario.InjetarRotas(mux)
	r.ManipuladorTarefa.InjetarRotas(mux)

	// INJETA INTERMEDIÁRIOS - Middlewares
	//r.IntermediarioAutenticador = intermediarios.NovoAutenticador(mux, r.ManipuladorAutenticacao.CDUAutenticacao)
	//r.IntermediarioLogger = intermediarios.NovoLogger(r.IntermediarioAutenticador)
	//r.Handler = r.IntermediarioLogger

	roteador := intermediarios.AutenticadorFunc(mux, r.ManipuladorAutenticacao.CDUAutenticacao)
	roteador = intermediarios.LoggerFunc(roteador)

	r.Handler = &roteador

}
