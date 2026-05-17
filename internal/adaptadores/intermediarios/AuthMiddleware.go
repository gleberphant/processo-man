package intermediarios

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/gleberphant/ProcessoMan/internal/casosdeuso/CasosDeUsoAutenticacao"
	"github.com/gleberphant/ProcessoMan/internal/entidades"
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

func AuthMiddleware(proximo http.Handler, auth *CasosDeUsoAutenticacao.CasosDeUsoAutenticacao) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//-------------------------------------
		// descatar pagina de login e fav icon
		// fav icon para nao dobrar a requisicao
		if r.URL.Path == "/favicon.ico" {
			w.WriteHeader(http.StatusNoContent) // Devolve 204 (Sem Conteúdo)
			return
		}

		// pagina de login nao precisa de autenticação, portanto, retorna para nao entrar em loop eterno
		if r.URL.Path == "/login" {
			proximo.ServeHTTP(w, r)
			return
		}

		//-------------------------------------
		// procurar token

		// 1º vamos tentar pegar  o token do cookie
		token, err := ProcurarTokenEnviado(r)

		if err != nil {
			log.Printf("Token não encontrado : [%v] ", err)
			http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso negado. Token não encontrado"), http.StatusSeeOther)
			return
		}

		//-------------------------------------
		// validar o token

		// 2º verificar se o token existe no  no banco de dados

		err = auth.ValidarToken(entidades.Token{UUID: uuid.MustParse(token)})

		// se houve erro na validação. Redireciona para LOGIN
		if err != nil {
			log.Printf("Token Inválido : [%v] ", err)
			http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso negado. Token Inválido"), http.StatusSeeOther)
			return
		}

		//se nao houve erro injeta token no cabeçalho então o acesso é permitido

		// iniciar seção com cookies
		proximo.ServeHTTP(w, r)
	})

}
