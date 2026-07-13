package roteamento

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/gleberphant/ProcessoMan/internal/intermediarios"
)

// configurar as rotas e devolver MUX configurado
func (r *Roteador) InjetarRotas() {

	// configurar mux
	mux := http.NewServeMux()

	// PAGINAS ESTÁTICAS
	// index
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		apresentacao.ExibirPaginaHTML("pages/index.html", w, r, nil)
	})

	// fallback para pagina 404
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		apresentacao.ExibirPaginaHTML("pages/page404.html", w, r, nil)
	})

	// Not found para o DevTool do google chrome
	mux.HandleFunc("/.well-known/*", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	// retornar o favicon da aplicação
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/favicon.ico") // Ajuste o caminho do seu arquivo
	})

	// ROTAS DOS MANIPULADORES
	r.ManipuladorAutenticacao.InjetarRotas(mux)
	r.ManipuladorProcesso.InjetarRotas(mux)
	r.ManipuladorUsuario.InjetarRotas(mux)
	r.ManipuladorTarefa.InjetarRotas(mux)

	// INJETA INTERMEDIÁRIOS - Middlewares
	roteador := intermediarios.AutenticadorFunc(mux, r.ManipuladorAutenticacao.CDUAutenticacao)
	roteador = intermediarios.LoggerFunc(roteador)

	r.Handler = &roteador

}
