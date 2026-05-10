package controladores

import (
	"html/template"
	"log"
	"net/http"
)

func FormularioLogin(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		http.Error(w, "Erro ao carregar template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	msg := r.URL.Query().Get("msg")

	// carregar dados
	msgPagina := struct {
		Msg string
	}{
		Msg: msg,
	}

	//executar template

	err = tmpl.Execute(w, msgPagina)
	if err != nil {
		log.Printf("erro ao executar template")
		http.Error(w, "Erro ao renderizar pagina", http.StatusInternalServerError)
		return
	}

}

func Logar(w http.ResponseWriter, r *http.Request) {
	//return
}

func Index(w http.ResponseWriter, r *http.Request) {
	//return

}
