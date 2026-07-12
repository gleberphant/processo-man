package apresentacao

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var cacheTemplates map[string]*template.Template

func CarregarTemplates() error {

	if cacheTemplates == nil {
		cacheTemplates = make(map[string]*template.Template)
	}

	baseDir := filepath.Join("..", "templates")

	// Mapeia funções para a interface
	mapaFuncoes := template.FuncMap{
		"formatarData": formatarData,
	}

	// Mapeia Layout
	arquivosLayout := []string{
		filepath.Join(baseDir, "_layout", "_layout.html"),
		filepath.Join(baseDir, "_layout", "_header.html"),
		filepath.Join(baseDir, "_layout", "_footer.html"),
		filepath.Join(baseDir, "_layout", "_navbar.html"),
	}

	// Mapeia paginas sem login
	caminhoLogin := filepath.Join(baseDir, "autenticacao", "login.html")

	// 3. Substituição do Glob pelo WalkDir (Recursividade Real e Segura)
	err := filepath.WalkDir(baseDir, func(caminho string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Ignora diretórios e arquivos que não são HTML
		if d.IsDir() || filepath.Ext(caminho) != ".html" {
			return nil
		}

		// Impede que os fragmentos de layout sejam mapeados como páginas renderizáveis
		if strings.Contains(caminho, "_layout") {
			return nil
		}

		tmpl := template.New(filepath.Base(caminho)).Funcs(mapaFuncoes)

		var arquivosAlvo []string
		if caminho != caminhoLogin { // Comparação agora funciona em qualquer SO
			arquivosAlvo = append(arquivosLayout, caminho)
		} else {
			arquivosAlvo = []string{caminho}
		}

		tmpl, err = tmpl.ParseFiles(arquivosAlvo...)
		if err != nil {
			log.Printf("erro ao compilar template %s . Error %v", caminho, err.Error())
			return err
		}

		// 4. Extração segura da chave (Relativo + ToSlash para compatibilidade Linux/Windows)
		chaveRelativa, err := filepath.Rel(baseDir, caminho)
		if err != nil {
			return err
		}

		chaveFinal := filepath.ToSlash(chaveRelativa)
		cacheTemplates[chaveFinal] = tmpl

		log.Printf("Chave %s Template %s", chaveFinal, caminho)
		return nil
	})

	return err
}
