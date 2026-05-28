package roteamento

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
)

//var layout = []string{"../templates/layout/_layout.html", "../templates/layout/_header.html", "../templates/layout/_navbar.html", "../templates/layout/_footer.html"}

func (s *Roteador) Index(w http.ResponseWriter, r *http.Request) {

	apresentacao.ExibirPaginaHTML("pages/index.html", w, nil)

}

func (s *Roteador) Pagina404(w http.ResponseWriter, r *http.Request) {

	apresentacao.ExibirPaginaHTML("pages/page1.html", w, nil)

}

func (s *Roteador) Page2(w http.ResponseWriter, r *http.Request) {

	apresentacao.ExibirPaginaHTML("pages/page2.html", w, nil)

}
