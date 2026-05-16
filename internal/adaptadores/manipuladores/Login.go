package manipuladores

import (
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/gleberphant/ProcessoMan/internal/casosdeuso/autenticacao"
	"github.com/gleberphant/ProcessoMan/internal/entidades"
)

type ManipuladorLogin struct {
	CDUatenticacao *autenticacao.AutenticacaoCDU
}

func NovoManipuladorLogin(cduAutenticacao *autenticacao.AutenticacaoCDU) *ManipuladorLogin {

	return &ManipuladorLogin{
		CDUatenticacao: cduAutenticacao,
	}

}

// formulario de login
func (m *ManipuladorLogin) LoginGet(w http.ResponseWriter, r *http.Request) {

	// carrega dados
	dados := struct {
		Msg string
	}{
		Msg: r.URL.Query().Get("msg"),
	}

	// carregat HTML
	tmpl, err := template.ParseFiles("../templates/login.html")
	if err != nil {
		http.Error(w, "Erro ao carregar template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// executar template
	err = tmpl.Execute(w, dados)
	if err != nil {
		http.Error(w, "Erro ao renderizar pagina", http.StatusInternalServerError)
		return
	}

}

// funcao para logar
func (m *ManipuladorLogin) LoginPost(w http.ResponseWriter, r *http.Request) {
	// pega os dados do login
	var usuario = entidades.Usuario{
		Email: r.PostFormValue("email"),
		Senha: r.PostFormValue("senha"),
	}

	// autenticacao do usuario
	token, err := m.CDUatenticacao.AutenticarUsuario(&usuario)

	if err != nil {
		log.Printf("Usuario invalido %v", err)
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
