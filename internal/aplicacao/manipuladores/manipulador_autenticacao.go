package manipuladores

import (
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominio/servicos"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
)

type ManipuladorAutenticacao struct {
	CDUAutenticacao *servicos.ServicoAutenticacao
}

func NovoManipuladorLogin(cduToken *servicos.ServicoAutenticacao) *ManipuladorAutenticacao {

	return &ManipuladorAutenticacao{
		CDUAutenticacao: cduToken,
	}

}

func (m *ManipuladorAutenticacao) Fechar() {
	m.CDUAutenticacao.Fechar()
}

// formulario de login
func (m *ManipuladorAutenticacao) PageLogin(w http.ResponseWriter, r *http.Request) {

	// carrega dados
	dados := struct {
		Msg string
	}{
		Msg: r.URL.Query().Get("msg"),
	}

	apresentacao.ExibirPaginaHTML("login.html", w, r, dados)

}

// funcao para logar
func (m *ManipuladorAutenticacao) LoginPost(w http.ResponseWriter, r *http.Request) {

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
