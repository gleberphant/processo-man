package autenticacao

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
)

type ManipuladorLogin struct {
	CDUAutenticacao *CDUAutenticacao
}

func NovoManipuladorLogin(cduToken *CDUAutenticacao) *ManipuladorLogin {

	return &ManipuladorLogin{
		CDUAutenticacao: cduToken,
	}

}

// formulario de login
func (m *ManipuladorLogin) PageLogin(w http.ResponseWriter, r *http.Request) {

	// carrega dados
	dados := struct {
		Msg string
	}{
		Msg: r.URL.Query().Get("msg"),
	}

	apresentacao.ExibirHTMLSemLayout("autenticacao/login.html", w, dados)

}

// funcao para logar
func (m *ManipuladorLogin) LoginPost(w http.ResponseWriter, r *http.Request) {
	// pega os dados do login
	var usuario = entidades.Usuario{
		Email: r.PostFormValue("email"),
		Senha: r.PostFormValue("senha"),
	}

	// autenticacao do usuario e retorna o token gerado
	token, err := m.CDUAutenticacao.AutenticarUsuario(usuario.Email, usuario.Senha)

	if err != nil {
		log.Printf("Erro na autenticacao. %v", err)
		http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso Negado. Usuario Inválido."), http.StatusSeeOther)
		return
	}

	// Configura o cookie de sessão de forma segura
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token.UUID.String(),
		//Path:     "/",
		MaxAge:   3600,                 // Define expiração para 1 hora (em segundos)
		HttpOnly: true,                 // Protege contra roubo via JavaScript (ataques XSS)
		SameSite: http.SameSiteLaxMode, // Protege contra requisições forjadas de outros sites (CSRF)
		// Secure: true,               // Descomente esta linha quando colocar em produção com HTTPS
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
