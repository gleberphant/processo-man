package controladores

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/modelos"
	"github.com/gleberphant/ProcessoMan/internal/servicos/casosdeuso"
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

// funcao para logar
func Logar(w http.ResponseWriter, r *http.Request) {

	// recebe a requisao
	// chama o servico/usecase login
	var usuario = modelos.Usuario{}

	usuario.Email = "root@root"
	usuario.Senha = "root"

	token, err := casosdeuso.Logar(usuario)

	if err != nil {
		log.Printf("Erro logar: %s", err)
		http.Error(w, "Erro ao logar no sistema", http.StatusInternalServerError)
		return
	}

	mapa := r.URL.Query()
	mapa.Set("token", token.Token)
	r.URL.RawQuery = mapa.Encode()

	Index(w, r)

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
