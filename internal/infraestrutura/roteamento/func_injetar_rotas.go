package roteamento

import (
	"net/http"
	"path/filepath"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
)

// configurar as rotas e devolver MUX configurado
func (r *Roteador) InjetarRotas() {

	// configurar mux
	mux := http.NewServeMux()

	// PAGINAS ESTÁTICAS
	// index
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, req *http.Request) {
		apresentacao.ExibirPaginaHTML("index.html", w, req, nil)
	})

	// fallback para pagina 404
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		apresentacao.ExibirPaginaHTML("page404.html", w, req, nil)
	})

	// Not found para o DevTool do google chrome
	mux.HandleFunc("/.well-known/*", func(w http.ResponseWriter, req *http.Request) {
		http.NotFound(w, req)
	})

	// retornar o favicon da aplicação
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, filepath.Join(r.appConfig.DiretorioRaiz, "/static/favicon.ico")) // Ajuste o caminho do seu arquivo
	})

	// INSERIR ROTAS DOS MANIPULADORES
	for _, m := range r.manipuladores {
		m.InjetarRotas(mux)
	}

	r.Handler = mux

}
