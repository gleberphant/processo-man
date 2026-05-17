package apresentacao

import (
	"html/template"
	"log"
	"net/http"
)

func ExibirPaginaHTML(page string, w http.ResponseWriter, dados interface{}) error {

	tmpl, err := template.ParseFiles(
		"../templates/_layout/_layout.html",
		"../templates/_layout/_header.html",
		"../templates/_layout/_navbar.html",
		"../templates/_layout/_footer.html",
		"../templates/"+page,
	)

	if err != nil {
		log.Printf("Erro ao carregar arquivos do template: %v", err)
		http.Error(w, "Erro ao carregar pagina", http.StatusInternalServerError)
		return err
	}

	err = tmpl.ExecuteTemplate(w, "_layout", dados)

	if err != nil {
		log.Printf("erro ao executar template: %v", err)
		http.Error(w, "Erro ao executar pagina", http.StatusInternalServerError)
		return err
	}

	return nil
}

func ExibirHTMLSemLayout(page string, w http.ResponseWriter, dados interface{}) error {

	tmpl, err := template.ParseFiles("../templates/" + page)

	if err != nil {
		log.Printf("erro ao carregar arquivos do template")
		http.Error(w, "Erro na renderização da pagina", http.StatusInternalServerError)
		return err
	}

	err = tmpl.Execute(w, dados)

	if err != nil {
		log.Printf("erro ao executar template")
		http.Error(w, "Erro na renderização da pagina", http.StatusInternalServerError)
		return err
	}

	return nil
}
