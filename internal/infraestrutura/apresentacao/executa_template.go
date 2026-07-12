package apresentacao

import (
	"log"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
)

func ExibirPaginaHTML(chave string, w http.ResponseWriter, r *http.Request, dados interface{}) error {

	// Busca o template pré-compilado do cache
	tmpl, ok := cacheTemplates[chave]
	if !ok {
		log.Printf("Erro ao carregar pagina: template não encontrado no cache %v", ok)
		http.Error(w, "Erro ao carregar pagina: template não encontrado no cache", http.StatusInternalServerError)
		return nil // ou um erro específico
	}

	// injetar dados globais da requisição

	var usuarioLogado string = "sem usuario"

	// Tenta obter o token do contexto da requisição.
	if tokenCtx := r.Context().Value("TokenContext"); tokenCtx != nil {

		// Faz a asserção de tipo para entidades.Token e verifica se foi bem-sucedida.
		if token, ok := tokenCtx.(entidades.Token); ok {

			usuarioLogado = token.UsuarioNome
		}

	}

	log.Printf("Usuario logado: %s", usuarioLogado)
	// Injetar dados globais no viewModel
	viewModeComContexto := struct {
		UsuarioLogado string
		Dados         interface{}
	}{
		// Se o token não for encontrado ou o tipo for inválido, UsuarioLogado será uma string vazia.
		UsuarioLogado: usuarioLogado,
		Dados:         dados,
	}

	// executa o template
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
