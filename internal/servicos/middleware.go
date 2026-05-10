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
		if r.URL.Path == "/favicon.ico" {
			w.WriteHeader(http.StatusNoContent) // Devolve 204 (Sem Conteúdo)
			return
		}

		// se for a pagina do login retorna
		if r.URL.Path == "/login" || r.URL.Path == "/favicon.ico" {
			proximo.ServeHTTP(w, r)
			return
		}

		// 1º vamos tentar pegar  o token do cookie ou do url get
		var token string //token inicia vazio

		// tenta ler o token por cookie
		cookie, err := r.Cookie("token_sessao")

		// se não tem erro no cookie. ler valor do cookie
		if err == nil {
			token = cookie.Value
		}
		token = "ABC"
		// se token do cookie vazio então procura no get
		if token == "" {
			token = r.URL.Query().Get("token")
		}

		// se token do GET também vazio, então retorna
		if token == "" {
			msg := url.QueryEscape("Token de acesso inválido. Acesso Negado.")
			http.Redirect(w, r, "/login?msg="+msg, http.StatusSeeOther)
			return
		}

		// 2º verificar se o token existe no  no banco de dados
		log.Print("Autenticando usuario: " + token)

		resultado, err := repositorio.Consultar("SELECT id FROM tokens WHERE token LIKE '?' ", token)

		// se houve erro na conexao com o  banco de dados. informar ao usuario
		if err != nil {
			msg := url.QueryEscape("Erro no servidor de Banco de Dados. Tente novamente")
			http.Redirect(w, r, "/login?msg="+msg, http.StatusSeeOther)
			return
		}

		defer resultado.Close()

		var acessoNegado bool = true

		if resultado.Next() {
			acessoNegado = false
		}

		// Verifica se acesso negado
		if acessoNegado {
			msg := url.QueryEscape("Login Inválido. Acesso Negado")
			http.Redirect(w, r, "/login?msg="+msg, http.StatusSeeOther)
			return
		}

		proximo.ServeHTTP(w, r)
	})

}
