package apresentacao

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func ExibirPaginaHTML(page string, w http.ResponseWriter, viewModel interface{}) error {

	var err error

	// cria novo template
	tmpl := template.New("pagina_html_com_layout")

	// Mapeia funções para a interface
	tmpl = tmpl.Funcs(template.FuncMap{
		"formatarData": formatarData, //função de formatação de formatação de data
	})

	// carrega os arquivos
	tmpl, err = tmpl.ParseFiles(
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

	//executa o template
	err = tmpl.ExecuteTemplate(w, "_layout", viewModel)

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

// 1. Crie a função de formatação de formatação de data
func formatarData(data time.Time) string {
	if data.IsZero() {
		return "Sem Data" // ou retorne "" se preferir vazio
	}
	// Formato padrão brasileiro
	return data.Format("02/01/2006")
}

func RedirecionarPaginaAnterior(w http.ResponseWriter, r *http.Request, fallback ...string) {

	destino := "/"

	if len(fallback) > 0 && fallback[0] != "" {
		destino = fallback[0]
	} else if referencia := r.Header.Get("Referer"); referencia != "" {
		destino = referencia
	}

	http.Redirect(w, r, destino, http.StatusSeeOther)
}
