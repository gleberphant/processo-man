package controladores

import (
	"html/template"
	"log"
	"net/http"
)

// pagina de login
func FormularioLogin(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("../templates/login.html")
	if err != nil {
		log.Printf("Erro ao carregar template: %v", err.Error())
		http.Error(w, "Erro ao carregar template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	msg := r.URL.Query().Get("msg")

	// carregar dados
	dados := struct {
		Msg string
	}{
		Msg: msg,
	}

	// executar template
	err = tmpl.Execute(w, dados)
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

	tmpl, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		log.Printf("Erro ao carregar template: %v", err.Error())
		http.Error(w, "Erro ao carregar template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// executar template
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("erro ao executar template")
		http.Error(w, "Erro ao renderizar pagina", http.StatusInternalServerError)
		return
	}

}
