package servicos

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gleberphant/ProcessoMan/internal/repositorio"
)

func LogMiddleware(proximo http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proximo.ServeHTTP(w, r)

		log.Printf("| Requisao: %s |  Metodo: %s | ",
			r.URL,
			r.Method,
		)
	})
}

func AuthMiddleware(proximo http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fav icon para nao dobrar a requisicao
		if r.URL.Path == "/favicon.ico" {
			w.WriteHeader(http.StatusNoContent) // Devolve 204 (Sem Conteúdo)
			return
		}

		// se for a pagina do login retorna para nao entrar em loop eterno
		if r.URL.Path == "/login" {
			proximo.ServeHTTP(w, r)
			return
		}

		// senao for a pagina de login então verifica as credenciais

		// 1º vamos tentar pegar  o token do cookie ou do url get
		var token string //token inicia vazio

		// tenta ler o token por cookie
		cookie, err := r.Cookie("token_sessao")

		// se não tem erro no cookie. ler valor do cookie
		if err == nil {
			token = cookie.Value
		}

		// se token do cookie é vazio então procura no get
		if token == "" {
			token = r.URL.Query().Get("token")
		}
		//token = "ABC"

		// se token do GET também está vazio, então retorna com mensagem
		if token == "" {
			rota := "/login?msg=" + url.QueryEscape("Acesso Negado. Forneca um TOKEN de acesso.")
			http.Redirect(w, r, rota, http.StatusSeeOther)
			return
		}

		// 2º verificar se o token existe no  no banco de dados
		log.Printf("Autenticando usuario com token: [%s] ", token)

		resultado, err := repositorio.Consultar("SELECT id FROM tokens WHERE token LIKE 'ABC'", token)

		// se houve erro na conexao com o  banco de dados. informar ao usuario
		if err != nil {
			rota := "/login?msg=" + url.QueryEscape("Erro no servidor de Banco de Dados. Tente novamente")
			http.Redirect(w, r, rota, http.StatusSeeOther)
			return
		}

		defer resultado.Close()

		var acessoNegado bool = true

		if resultado.Next() {
			log.Printf("Usuario autenticado com token: %s ", token)
			acessoNegado = false
		}

		// Verifica se acesso negado
		if acessoNegado {
			rota := "/login?msg=" + url.QueryEscape("Acesso negado. token não encontrado")
			http.Redirect(w, r, rota, http.StatusSeeOther)
			return
		}

		proximo.ServeHTTP(w, r)
	})

}
