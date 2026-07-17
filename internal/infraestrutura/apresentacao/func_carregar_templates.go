package apresentacao

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var cacheTemplates map[string]*template.Template

func CarregarTemplates(diretorioRaiz string) error {

	if cacheTemplates == nil {
		cacheTemplates = make(map[string]*template.Template)
	}

	diretorioTemplates := filepath.Join(diretorioRaiz, "templates")

	// Mapeia funções para a interface
	mapaFuncoes := template.FuncMap{
		"formatarData": formatarData,
	}

	// Mapeia Layout
	arquivosLayout := []string{
		filepath.Join(diretorioTemplates, "_layout", "_layout.html"),
		filepath.Join(diretorioTemplates, "_layout", "_header.html"),
		filepath.Join(diretorioTemplates, "_layout", "_footer.html"),
		filepath.Join(diretorioTemplates, "_layout", "_navbar.html"),
	}

	// 1. Encontrar todos os arquivos de componentes reutilizáveis
	var arquivosComponentes []string
	err := filepath.WalkDir(diretorioTemplates, func(caminho string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Adiciona arquivos que estão dentro de um diretório chamado "componentes"
		if !d.IsDir() && strings.Contains(filepath.ToSlash(caminho), "/componentes/") {
			arquivosComponentes = append(arquivosComponentes, caminho)
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Mapeia pagina login sem layout
	caminhoLogin := filepath.Join(diretorioTemplates, "autenticacao", "login.html")

	// 2. Percorrer novamente para criar as páginas, agora injetando os componentes
	err = filepath.WalkDir(diretorioTemplates, func(caminho string, baseName os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Ignora diretórios e arquivos que não são HTML
		if baseName.IsDir() {
			return nil
		}

		// Impede que os fragmentos de layout sejam mapeados como páginas renderizáveis
		if strings.Contains(filepath.ToSlash(caminho), "/_layout/") || strings.Contains(filepath.ToSlash(caminho), "/componentes/") {
			return nil
		}

		// cria um novo template com o mapa de funcoes
		tmpl := template.New(filepath.Base(caminho)).Funcs(mapaFuncoes)

		// Junta os arquivos de layout, a página atual e todos os componentes
		var arquivosParaParse []string
		if caminho != caminhoLogin { // Comparação agora funciona em qualquer SO
			arquivosParaParse = append(arquivosLayout, caminho)
			arquivosParaParse = append(arquivosParaParse, arquivosComponentes...)
		} else {
			// A página de login não usa layout nem componentes
			arquivosParaParse = []string{caminho}
		}

		tmpl, err = tmpl.ParseFiles(arquivosParaParse...)
		if err != nil {
			log.Printf("erro ao compilar template %s . Error %v", caminho, err.Error())
			return err
		}

		// Extração segura da chave (Relativo + ToSlash para compatibilidade Linux/Windows)
		chaveRelativa, err := filepath.Rel(diretorioTemplates, caminho)
		if err != nil {
			return err
		}

		chaveFinal := filepath.ToSlash(chaveRelativa)
		cacheTemplates[chaveFinal] = tmpl

		//log.Printf("Chave %s Template %s", chaveFinal, caminho)
		return nil
	})
	log.Printf("Templates [%d] Carregados ", len(cacheTemplates))

	return err
}
