package roteamento

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
)

//var layout = []string{"../templates/layout/_layout.html", "../templates/layout/_header.html", "../templates/layout/_navbar.html", "../templates/layout/_footer.html"}

func (r *Roteador) Index(w http.ResponseWriter, req *http.Request) {

	apresentacao.ExibirPaginaHTML("pages/index.html", w, nil)

}

func (r *Roteador) Pagina404(w http.ResponseWriter, req *http.Request) {

	apresentacao.ExibirPaginaHTML("pages/page1.html", w, nil)

}

func (r *Roteador) Page2(w http.ResponseWriter, req *http.Request) {

	apresentacao.ExibirPaginaHTML("pages/page2.html", w, nil)

}
