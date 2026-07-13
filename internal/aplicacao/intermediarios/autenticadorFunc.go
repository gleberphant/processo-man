package intermediarios

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gleberphant/ProcessoMan/internal/dominio/servicos"
	"github.com/google/uuid"
)

func AutenticadorFunc(proximo http.Handler, servicoAutenticacao *servicos.ServicoAutenticacao) http.Handler {

	rotasLivres := map[string]bool{
		"/login":       true,
		"/favicon.ico": true,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//-------------------------------------
		// bypass de login e fav icon
		if rotasLivres[r.URL.Path] || strings.HasPrefix(r.URL.Path, "/.well-known/") {
			log.Printf(" - byPass autenticacao: %s", r.URL.Path)
			proximo.ServeHTTP(w, r)
			return
		}

		//log.Printf(" - Realizando autenticação  : %s", r.URL.Path)
		//-------------------------------------
		// Extrai o token enviado
		strTokenUUID, err := procurarTokenEnviado(r)

		if err != nil {
			log.Printf("Token não encontrado : [%v] ", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		//-------------------------------------
		// Converte o token para formato UUID
		tokenUUID, err := uuid.Parse(strTokenUUID)
		if err != nil {
			log.Printf("Formato de token inválido: [%v] ", err)
			http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso negado. Token Inválido"), http.StatusSeeOther)
			return
		}

		//Verifica se o token existe
		token, err := servicoAutenticacao.VerificarExisteToken(tokenUUID)
		if err != nil {
			log.Printf("Token inexistente: [%v] ", err)
			http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso negado. Token Expirado"), http.StatusSeeOther)
			return
		}

		// Verifica a permissão do token
		// Extrai apenas a rota principal da URL -- implementação provisoria
		partes := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		rota := "/" + partes[0]

		err = servicoAutenticacao.VerificarPermissao(token, rota, r.Method)

		if err != nil {
			log.Printf("Erro Permissao: [%v] ", err)
			http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso negado. Usuario sem permissão"), http.StatusSeeOther)
			return
		}

		// token autorizado. injeta usuario.
		ctxAutenticacao := context.WithValue(r.Context(), "TokenContext", *token)

		// iniciar seção
		proximo.ServeHTTP(w, r.WithContext(ctxAutenticacao))
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
