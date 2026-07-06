package apresentacao

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

var cacheTemplates map[string]*template.Template

func CarregarTemplates() error {

	if cacheTemplates == nil {
		cacheTemplates = make(map[string]*template.Template)
	}

	// buscar todos arquivos que começam com a palavra page

	paginas, err := filepath.Glob("../templates/**/*.html")

	if err != nil {
		return err
	}

	for _, page := range paginas {

		// cria novo template
		tmpl := template.New(filepath.Base(page))

		// Mapeia funções para a interface
		tmpl = tmpl.Funcs(template.FuncMap{
			"formatarData": formatarData, //função de formatação de formatação de data
		})

		if page == "../templates/autenticacao/login.html" {
			tmpl, err = tmpl.ParseFiles(page)
		} else {

			tmpl, err = tmpl.ParseFiles(
				"../templates/_layout/_layout.html",
				"../templates/_layout/_header.html",
				"../templates/_layout/_footer.html",
				"../templates/menu/_navbar.html",
				page,
			)
		}

		if err != nil {
			log.Printf("Erro ao carregar arquivos do templates: %v", err)
			return err
		}

		// 1. Remove o "..\templates\" do início
		chave := strings.TrimPrefix(page, `..\templates\`)

		// 2. Substitui as barras invertidas (\) por barras normais (/)
		chave = strings.ReplaceAll(chave, `\`, `/`)
		cacheTemplates[chave] = tmpl

		log.Printf("Template carregado e cacheado: %s", page)

	}
	return nil
}

func ExibirPaginaHTML(chave string, w http.ResponseWriter, viewModel interface{}) error {

	// Busca o template pré-compilado do cache.
	tmpl, ok := cacheTemplates[chave]
	if !ok {
		log.Printf("Erro ao carregar pagina: template não encontrado no cache %v", ok)
		http.Error(w, "Erro ao carregar pagina: template não encontrado no cache", http.StatusInternalServerError)
		return nil // ou um erro específico
	}

	//executa o template
	err := tmpl.ExecuteTemplate(w, "_layout", viewModel)

	if err != nil {
		log.Printf("erro ao executar template: %v", err)
		http.Error(w, "Erro ao executar pagina", http.StatusInternalServerError)
		return err
	}

	return nil
}

func ExibirHTMLSemLayout(chave string, w http.ResponseWriter, viewModel interface{}) error {

	// Busca o template pré-compilado do cache.
	tmpl, ok := cacheTemplates[chave]
	if !ok {
		log.Printf("Erro ao carregar pagina sem layout: template não encontrado no cache %v", ok)
		http.Error(w, "Erro ao carregar pagina sem layout: template não encontrado no cache", http.StatusInternalServerError)
		return nil
	}

	err := tmpl.Execute(w, viewModel)

	if err != nil {
		log.Printf("erro ao executar template sem layout %v", err)
		http.Error(w, "Erro na renderização da pagina sem layout", http.StatusInternalServerError)
		return err
	}

	return nil
}
