package apresentacao

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func ExibirErro(w http.ResponseWriter, erroMsg string) {

	log.Println(erroMsg)
	//substituir por redirecionamento para o index com uma mensagem
	http.Error(w, erroMsg, http.StatusInternalServerError)

}

func ExibirJsonApi(w http.ResponseWriter, dados interface{}) error {
	// Converte  'dados' para o formato JSON.
	jason, err := json.Marshal(dados)

	if err != nil {
		return err
	}

	// Define o cabeçalho da resposta para indicar que o conteúdo é JSON.
	w.Header().Set("Contet-Type", "application/json")

	// Escreve o JSON resultante no corpo da resposta.
	_, err = w.Write(jason)

	if err != nil {
		return err
	}
	return nil
}

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
