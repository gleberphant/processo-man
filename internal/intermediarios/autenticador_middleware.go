package intermediarios

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
	"github.com/google/uuid"
)

func ProcurarTokenEnviado(r *http.Request) (string, error) {
	var token string //token inicia vazio

	// // tenta ler o token por cookie
	cookie, err := r.Cookie("token")

	// se não tem erro no cookie. ler valor do cookie
	if err == nil {
		token = cookie.Value
	}

	// se token do cookie é vazio então procura no get
	if token == "" {
		token = r.URL.Query().Get("token")
	}

	// se token do GET também está vazio, então retorna com mensagem
	if token == "" {
		return "", errors.New("Token não encontrado")
	}

	return token, nil
}

func AutenticadorIntermediario(proximo http.Handler, autenticador *autenticacao.CDUAutenticacao) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//-------------------------------------
		// DESCARTA REQUEST DE login e fav icon
		if r.URL.Path == "/favicon.ico" {
			w.WriteHeader(http.StatusNoContent) // Devolve 204 (Sem Conteúdo)
			return
		}

		if r.URL.Path == "/login" {
			proximo.ServeHTTP(w, r)
			return
		}

		//-------------------------------------
		// procurar token
		token, err := ProcurarTokenEnviado(r)

		if err != nil {
			log.Printf("Token não encontrado : [%v] ", err)
			http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso negado. Token não encontrado"), http.StatusSeeOther)
			return
		}

		//-------------------------------------
		// validar o token
		tokenUUID, err := uuid.Parse(token)
		if err != nil {
			log.Printf("Formato de token inválido: [%v] ", err)
			http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso negado. Token Inválido"), http.StatusSeeOther)
			return
		}

		err = autenticador.ValidarToken(tokenUUID)

		// se houver erro na validação. Redireciona para LOGIN
		if err != nil {
			log.Printf("Token Inválido : [%v] ", err)
			http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso negado. Token Inválido"), http.StatusSeeOther)
			return
		}

		// iniciar seção
		proximo.ServeHTTP(w, r)
	})

}
