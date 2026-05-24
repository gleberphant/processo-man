package autenticacao

import (
	"fmt"
	"net/http"

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
	viewModel := struct {
		Msg string
	}{
		Msg: r.URL.Query().Get("msg"),
	}

	apresentacao.ExibirHTMLSemLayout("autenticacao/login.html", w, viewModel)

}

// funcao para logar
func (m *ManipuladorLogin) LoginPost(w http.ResponseWriter, r *http.Request) {

	// autenticacao do usuario e retorna o token gerado
	tokenUUID, err := m.CDUAutenticacao.AutenticarUsuario(r.PostFormValue("email"), r.PostFormValue("senha"))

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Login Post: %v", err))
		return
	}

	// Configura o cookie de sessão de forma segura
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenUUID,
		//Path:     "/",
		MaxAge:   3600,                 // Define expiração para 1 hora (em segundos)
		HttpOnly: true,                 // Protege contra roubo via JavaScript (ataques XSS)
		SameSite: http.SameSiteLaxMode, // Protege contra requisições forjadas de outros sites (CSRF)
		// Secure: true,               // Descomente esta linha quando colocar em produção com HTTPS
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
