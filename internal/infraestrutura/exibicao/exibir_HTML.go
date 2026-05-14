package exibicao

import (
	"html/template"
	"log"
	"net/http"
)

func ExibirPaginaHTML(page string, w http.ResponseWriter, dados interface{}) error {

	var layouts = []string{"../templates/layout/_layout.html", "../templates/layout/_header.html", "../templates/layout/_navbar.html", "../templates/layout/_footer.html"}

	tmpl, err := template.ParseFiles(append(layouts, "../templates/"+page)...)

	if err != nil {
		log.Printf("erro ao carregar arquivos do template")
		http.Error(w, "Erro na renderização da pagina", http.StatusInternalServerError)
		return err
	}

	err = tmpl.ExecuteTemplate(w, "_layout", dados)

	if err != nil {
		log.Printf("erro ao executar template")
		http.Error(w, "Erro na renderização da pagina", http.StatusInternalServerError)
		return err
	}

	return nil
}
