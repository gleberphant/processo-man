package manipuladores

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
)

//var layout = []string{"../templates/layout/_layout.html", "../templates/layout/_header.html", "../templates/layout/_navbar.html", "../templates/layout/_footer.html"}

func Index(w http.ResponseWriter, req *http.Request) {
	apresentacao.ExibirPaginaHTML("pages/index.html", w, req, nil)
}

func Pagina404(w http.ResponseWriter, req *http.Request) {
	apresentacao.ExibirPaginaHTML("pages/page404.html", w, req, nil)
}
