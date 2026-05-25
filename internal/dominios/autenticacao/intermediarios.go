package autenticacao

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

func AutenticadorIntermediario(proximo http.Handler, autenticador *CDUAutenticacao) http.Handler {

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
		strTokenUUID, err := procurarTokenEnviado(r)

		if err != nil {
			log.Printf("Token não encontrado : [%v] ", err)
			http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso negado. Token não encontrado"), http.StatusSeeOther)
			return
		}

		//-------------------------------------
		// validar o token
		tokenUUID, err := uuid.Parse(strTokenUUID)
		if err != nil {
			log.Printf("Formato de token inválido: [%v] ", err)
			http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso negado. Token Inválido"), http.StatusSeeOther)
			return
		}
		log.Printf("\n -------------- Verificando permissao ----------------")

		// Extrai apenas o recurso principal da URL
		caminhoLimpo := strings.Trim(r.URL.Path, "/")
		partes := strings.Split(caminhoLimpo, "/")
		recursoBase := "/" + partes[0]
		log.Printf("\n  uri : [%s]", recursoBase)

		err = autenticador.VerificarPermissao(tokenUUID, recursoBase)

		// se houver erro na validação. Redireciona para LOGIN
		if err != nil {
			log.Printf("Erro Permissão: token [%s] : Erro  [%v] ", tokenUUID, err)
			http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso negado <br> Usuario sem permissão"), http.StatusSeeOther)
			return
		}
		log.Printf("\n -------------- fim verificacao ----------------")
		// iniciar seção
		proximo.ServeHTTP(w, r)
	})

}

func LogIntermediario(proximo http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proximo.ServeHTTP(w, r)

		log.Printf("| Requisao: %s |  Metodo: %s | ",
			r.URL,
			r.Method,
		)
	})
}

func procurarTokenEnviado(r *http.Request) (string, error) {
	var token string //token inicia vazio

	//procura token no cabeçalho. padrão de API
	token = r.Header.Get("Authorization")
	if token != "" {
		token = strings.TrimPrefix(token, "Bearer")
	}

	// se token vazio procura token na url
	if token == "" {
		token = r.URL.Query().Get("token")
	}

	// se token  vazio procura cookie
	if token == "" {

		cookie, err := r.Cookie("token")

		if err == nil {
			token = cookie.Value
		}
	}

	// se token do cookie continua  vazio, então retorna com mensagem
	if token == "" {
		return "", errors.New(" Token não encontrado ")
	}

	return token, nil
}
