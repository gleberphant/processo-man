package apresentacao

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// cacheTemplates é um mapa que armazena os templates pré-compilados.
// A chave é o caminho relativo do template (ex: "usuario/index.html") e o valor é o objeto de template compilado.
// Usar um cache evita a releitura e compilação dos arquivos do disco a cada requisição, melhorando o desempenho.
var cacheTemplates map[string]*template.Template

// CarregarTemplates percorre o diretório de templates, compila cada página com seus layouts e componentes,
// e armazena o resultado em um cache na memória.
func CarregarTemplates(diretorioRaiz string) error {

	// Inicializa o mapa de cache se ele ainda não existir (lazy initialization).
	if cacheTemplates == nil {
		cacheTemplates = make(map[string]*template.Template)
	}

	// Constrói o caminho completo para o diretório de templates de forma segura entre sistemas operacionais.
	diretorioLayout := filepath.Join(diretorioRaiz, "templates", "layout")
	diretorioComponentes := filepath.Join(diretorioRaiz, "templates", "componentes")
	diretorioPaginas := filepath.Join(diretorioRaiz, "templates", "paginas")
	diretorioStatic := filepath.Join(diretorioRaiz, "static")

	// mapaFuncoes define um mapa de funções Go que podem ser chamadas de dentro dos templates.
	// Aqui, a função Go `formatarData` estará disponível no template como `{{ formatarData .AlgumaData }}`.
	mapaFuncoes := template.FuncMap{
		"formatarData": formatarData,
	}

	// arquivosLayout define a lista de arquivos que compõem a estrutura base de todas as páginas (exceto exceções).
	var err error

	// carrega a lista de arquivos de layout padrão
	var listaLayout []string
	err = filepath.WalkDir(diretorioLayout, func(filePath string, fileInfo os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() || filepath.Ext(filePath) != ".tmpl" {
			return nil
		}

		listaLayout = append(listaLayout, filePath)

		return nil

	})

	//  Carrega a lista de componentes  reutilizáveis. que podem ser incluídos em várias páginas.
	var listaComponentes []string
	err = filepath.WalkDir(diretorioComponentes, func(filePath string, fileInfo os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if fileInfo.IsDir() || filepath.Ext(filePath) != ".tmpl" {
			return nil
		}

		listaComponentes = append(listaComponentes, filePath)

		return nil
	})

	//  Percorre o diretorio de Páginas para gerar um template para cada arquivo .html encontrado
	err = filepath.WalkDir(diretorioPaginas, func(arquivoCaminho string, arquivoInfo os.DirEntry, err error) error {
		// se houver erro retorna
		if err != nil {
			return err
		}

		// ignora arquivos com extensões diferente de html ou tmpl
		if arquivoInfo.IsDir() || filepath.Ext(arquivoCaminho) != ".html" {
			return nil
		}

		// redundancia para verificar se o arquivo é de layout ou componente,
		if strings.Contains(arquivoCaminho, "layout") || strings.Contains(arquivoCaminho, "componente") {
			return nil
		}

		// Prepara a lista de arquivos que serão compilados juntos para formar o template
		listaArquivosDoTemplate := append(listaLayout, arquivoCaminho)
		listaArquivosDoTemplate = append(listaArquivosDoTemplate, listaComponentes...)

		// Gera uma chave única para o cache, baseada no caminho relativo do arquivo em relação ao diretório `paginas`.
		// Ex: "g:\...\templates\paginas\usuario\index.html" se torna "usuario/index.html".
		chave, err := filepath.Rel(diretorioPaginas, arquivoCaminho)
		if err != nil {
			return err
		}

		// Garante que a chave use barras normais `/` como separador, para consistência.
		chave = filepath.ToSlash(chave)

		// Cria um novo template, com o nome base do arquivo e o mapa de funções customizadas.
		tmpl := template.New(chave).Funcs(mapaFuncoes)

		// Compila o template (faz o "parse") com os arquivos listados. O resultado é um único template compilado.
		tmpl, err = tmpl.ParseFiles(listaArquivosDoTemplate...)
		if err != nil {
			log.Printf("erro ao compilar template %s . Error %v", arquivoCaminho, err.Error())
			return err
		}

		// Armazena o template compilado no cache com a chave gerada.
		cacheTemplates[chave] = tmpl
		return nil
	})

	// percorre o diretorio de Paginas Estaticas. para gerar um template sem layout para cada html encontrado
	err = filepath.WalkDir(diretorioStatic, func(arquivoCaminho string, arquivoInfo os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if arquivoInfo.IsDir() || filepath.Ext(arquivoCaminho) != ".html" {
			return nil
		}

		// Gera uma chave única para o cache, baseada no caminho relativo do arquivo em relação ao diretório `paginas`.
		// Ex: "g:\...\templates\paginas\usuario\index.html" se torna "usuario/index.html".
		chave, err := filepath.Rel(diretorioStatic, arquivoCaminho)
		if err != nil {
			return err
		}

		// Garante que a chave use barras normais `/` como separador, para consistência.
		chave = filepath.ToSlash(chave)

		// Cria um novo template, com o nome base do arquivo e o mapa de funções customizadas.
		tmpl := template.New(chave).Funcs(mapaFuncoes)

		// Compila o template (faz o "parse") apenas com o arquivo da pagina, sem outros arquivos.
		tmpl, err = tmpl.ParseFiles(arquivoCaminho)
		if err != nil {
			log.Printf("erro ao compilar template %s . Error %v", arquivoCaminho, err.Error())
			return err
		}

		// Armazena o template compilado no cache com a chave gerada.
		cacheTemplates[chave] = tmpl
		return nil
	})

	// Faz log da quantidade de templates carregados para confirmar que o processo foi executado.
	log.Printf("Templates [%d] Carregados ", len(cacheTemplates))

	return err
}
