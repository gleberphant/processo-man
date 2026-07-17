package apresentacao

import (
	"log"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
)

func ExibirPaginaHTML(chave string, w http.ResponseWriter, r *http.Request, dados interface{}) error {

	// Busca o template pré-compilado do cache
	tmpl, ok := cacheTemplates[chave]
	if !ok {
		log.Printf("Template [%s] não encontrado no cache %v", chave, ok)
		http.Error(w, "Template não encontrado no cache", http.StatusInternalServerError)
		return nil // ou um erro específico
	}

	// injetar dados globais da requisição
	var token entidades.Token = entidades.Token{}

	// Tenta obter o token do contexto da requisição.
	if tokenCtx := r.Context().Value("TokenContext"); tokenCtx != nil {

		// Faz a asserção de tipo para entidades.Token e verifica se foi bem-sucedida.
		token, _ = tokenCtx.(entidades.Token)

	}

	log.Printf("Usuario logado: %s", token.UsuarioNome)
	// Injetar dados globais no viewModel

	viewModel := struct {
		Menu entidades.Token
		Base interface{}
	}{
		// Se o token não for encontrado ou o tipo for inválido, UsuarioLogado será uma string vazia.
		Menu: token,
		Base: dados,
	}

	var err error
	if chave == "login/login.html" {
		err = tmpl.ExecuteTemplate(w, "login", dados)
	} else {
		err = tmpl.ExecuteTemplate(w, "_layout", viewModel)
	}
	// executa o template

	if err != nil {
		log.Printf("erro ao executar template: %v", err)
		http.Error(w, "Erro ao executar pagina", http.StatusInternalServerError)
		return err
	}

	return nil
}
