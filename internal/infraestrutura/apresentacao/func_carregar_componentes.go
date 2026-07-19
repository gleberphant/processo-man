package apresentacao

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// cachecomponentes é um mapa que armazena os componentes pré-compilados.
// A chave é o caminho relativo do componente (ex: "usuario/index.html") e o valor é o objeto de template compilado.
// Usar um cache evita a releitura e compilação dos arquivos do disco a cada requisição, melhorando o desempenho.
// esse cache unico evitar carregar todos componentes em cada pagina

var cacheComponentes map[string]*template.Template

// função para carregar cache componente
func CarregarComponentes(diretorioRaiz string, mapaFuncoes template.FuncMap) error {

	cacheComponentes = make(map[string]*template.Template)

	diretorioComponentes := filepath.Join(diretorioRaiz, "templates", "componentes")

	err := filepath.WalkDir(diretorioComponentes, func(arquivoCaminho string, arquivoInfo fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if arquivoInfo.IsDir() || filepath.Ext(arquivoCaminho) != ".tmpl" {
			return nil
		}

		chave, err := filepath.Rel(diretorioComponentes, arquivoCaminho)
		if err != nil {
			return err
		}

		chave, _ = strings.CutSuffix(chave, ".tmpl")
		arquivo, _ := os.ReadFile(arquivoCaminho)

		tmpl := template.New(chave).Funcs(mapaFuncoes)
		tmpl, err = tmpl.Parse(string(arquivo))

		if err != nil {
			return fmt.Errorf("falha ao analisar o template do componente %s: %w", arquivoCaminho, err)
		}

		cacheComponentes[chave] = tmpl

		return nil
	})

	return err
}
