package paginasestaticas

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
)

//var layout = []string{"../templates/layout/_layout.html", "../templates/layout/_header.html", "../templates/layout/_navbar.html", "../templates/layout/_footer.html"}

func Index(w http.ResponseWriter, r *http.Request) {

	apresentacao.ExibirPaginaHTML("index.html", w, nil)

}

func Page1(w http.ResponseWriter, r *http.Request) {

	apresentacao.ExibirPaginaHTML("page1.html", w, nil)

}

func Page2(w http.ResponseWriter, r *http.Request) {

	apresentacao.ExibirPaginaHTML("page2.html", w, nil)

}
