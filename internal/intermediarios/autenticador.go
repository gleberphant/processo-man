package intermediarios

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
	"github.com/google/uuid"
)

type Autenticador struct {
	servicoAutenticacao *autenticacao.CDUAutenticacao
	proximo             http.Handler
	mapaRotasLivres     map[string]bool
}

// construtor
func NovoAutenticador(proximoHandler http.Handler, servicoAutenticacao *autenticacao.CDUAutenticacao) *Autenticador {

	rotasLivres := map[string]bool{
		"/login":       true,
		"/favicon.ico": true,
	}

	return &Autenticador{
		proximo:             proximoHandler,
		mapaRotasLivres:     rotasLivres,
		servicoAutenticacao: servicoAutenticacao,
	}

}

func (a *Autenticador) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// descarta rotas sem autenticação. chama proximo handler
	if a.mapaRotasLivres[req.URL.Path] {
		a.proximo.ServeHTTP(res, req)
		return
	}

	//-------------------------------------
	// extrai o UUID do token enviado
	strTokenUUID, err := a.procurarTokenEnviado(req)

	// token não encontrado direciona para o login
	if err != nil {
		log.Printf("Token não encontrado : [%v] ", err)
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}

	// UUID UUID do token  encontrado , converte para formato UUID
	tokenUUID, err := uuid.Parse(strTokenUUID)
	if err != nil {
		log.Printf("Formato de token inválido: [%v] ", err)
		http.Redirect(res, req, "/login?msg="+url.QueryEscape("Token de acesso inválido"), http.StatusSeeOther)
		return
	}

	// Buscar o token, verifica se o token existe
	token, err := a.servicoAutenticacao.VerificarExisteToken(tokenUUID)
	if err != nil {
		log.Printf("Token Inexistente: [%v] ", err)
		http.Redirect(res, req, "/login?msg="+url.QueryEscape("Token Expirado . Acesso Negado"), http.StatusSeeOther)
		return
	}

	// Token encontrado(existe), verificar a permissão do token
	// Extrai apenas a rota principal da URL -- implementação provisoria
	partes := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	rota := "/" + partes[0]

	err = a.servicoAutenticacao.VerificarPermissao(token, rota, req.Method)
	if err != nil {
		log.Printf("Erro Permissao: [%v] ", err)
		http.Redirect(res, req, "/login?msg="+url.QueryEscape("Acesso negado. Usuario sem permissão"), http.StatusSeeOther)
		return
	}

	// token autorizado. injeta usuario no contexto.
	//ctxAutenticacao := context.WithValue(req.Context(), "Token", token)

	// chama proximo Handler
	log.Printf("Usuario autenticado: [%s] ", token.UsuarioNome)
	a.proximo.ServeHTTP(res, req /*req.WithContext(ctxAutenticacao)*/)

}

func (a *Autenticador) procurarTokenEnviado(r *http.Request) (string, error) {
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
