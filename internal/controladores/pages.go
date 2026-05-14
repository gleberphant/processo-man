package controladores

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/exibicao"
)

//var layout = []string{"../templates/layout/_layout.html", "../templates/layout/_header.html", "../templates/layout/_navbar.html", "../templates/layout/_footer.html"}

func Index(w http.ResponseWriter, r *http.Request) {

	exibicao.ExibirPaginaHTML("index.html", w, nil)

}

func Page1(w http.ResponseWriter, r *http.Request) {

	exibicao.ExibirPaginaHTML("page1.html", w, nil)

}

func Page2(w http.ResponseWriter, r *http.Request) {

	exibicao.ExibirPaginaHTML("page2.html", w, nil)

}
