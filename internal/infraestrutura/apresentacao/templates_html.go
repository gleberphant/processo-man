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
			"formatarData":  formatarData,
			"usuarioLogado": usuarioLogado, //função de formatação de formatação de data
		}

		tmpl = tmpl.Funcs(mapaFuncoes)

		layout := []string{filepathPagina}

		if filepathPagina != "../templates/autenticacao/login.html" {
			layout = append(layout,
				"../templates/_layout/_layout.html",
				"../templates/_layout/_header.html",
				"../templates/_layout/_footer.html",
				"../templates/_layout/_navbar.html",
			)
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

		log.Printf("Template carregado e cacheado: %s", filepathPagina)

	}
	return nil
}

func ExibirPaginaHTML(chave string, w http.ResponseWriter, r *http.Request, dados interface{}) error {
	// injetar dados globais da requisição

	// Busca o template pré-compilado do cache.

	tmpl, ok := cacheTemplates[chave]
	if !ok {
		log.Printf("Erro ao carregar pagina: template não encontrado no cache %v", ok)
		http.Error(w, "Erro ao carregar pagina: template não encontrado no cache", http.StatusInternalServerError)
		return nil // ou um erro específico
	}

	// injetar dados globais no viewModel

	viewModeComContexto := struct {
		UsuarioLogado string
		Dados         interface{}
	}{
		UsuarioLogado: "Usuario",
		Dados:         dados,
	}

	//executa o template
	err := tmpl.ExecuteTemplate(w, "_layout", viewModeComContexto)

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
