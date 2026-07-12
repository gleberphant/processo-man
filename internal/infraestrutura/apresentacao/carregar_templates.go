package apresentacao

import (
	"html/template"
	"log"
	"path/filepath"
	"strings"
)

var cacheTemplates map[string]*template.Template

func CarregarTemplates() error {

	if cacheTemplates == nil {
		cacheTemplates = make(map[string]*template.Template)
	}

	// buscar todos arquivos em templates que terminam com html
	listaFilepathPaginas, err := filepath.Glob("../templates/**/*.html")

	if err != nil {
		return err
	}

	for _, filepathPagina := range listaFilepathPaginas {
		// cria novo template
		tmpl := template.New(filepath.Base(filepathPagina))

		// Mapeia funções para a interface
		mapaFuncoes := template.FuncMap{
			"formatarData": formatarData,
		}

		tmpl = tmpl.Funcs(mapaFuncoes)

		var layout []string

		if filepathPagina != "../templates/autenticacao/login.html" {
			layout = []string{
				"../templates/_layout/_layout.html",
				"../templates/_layout/_header.html",
				"../templates/_layout/_footer.html",
				"../templates/_layout/_navbar.html",
				filepathPagina,
			}
		} else {
			layout = []string{filepathPagina}
		}

		tmpl, err = tmpl.ParseFiles(layout...)

		if err != nil {
			log.Printf("Erro ao carregar arquivos do templates: %v", err)
			return err
		}

		// 1. Remove o "..\templates\" do início
		chave := strings.TrimPrefix(filepathPagina, `..\templates\`)

		// 2. Substitui as barras invertidas (\) por barras normais (/)
		chave = strings.ReplaceAll(chave, `\`, `/`)
		cacheTemplates[chave] = tmpl

		log.Printf("Chave: %s , Template: %s", chave, filepathPagina)

	}
	return nil
}
