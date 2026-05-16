package manipuladores

import (
	"html/template"
	"net/http"
	"net/url"

	"github.com/gleberphant/ProcessoMan/internal/casosdeuso/usuarios"
	"github.com/gleberphant/ProcessoMan/internal/modelos"
)

// formulario de login
func LoginGet(w http.ResponseWriter, r *http.Request) {

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
func LoginPost(w http.ResponseWriter, r *http.Request) {
	// pega os dados do login
	var usuario = modelos.Usuario{
		Email: r.PostFormValue("email"),
		Senha: r.PostFormValue("senha"),
	}

	// autenticacao do usuario
	token, err := usuarios.AutenticarUsuario(&usuario)

	if err != nil {
		http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso Negado."), http.StatusSeeOther)
		return
	}

	// Configura o cookie de sessão de forma segura
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token.UUID,
		//Path:     "/",
		MaxAge:   3600,                 // Define expiração para 1 hora (em segundos)
		HttpOnly: true,                 // Protege contra roubo via JavaScript (ataques XSS)
		SameSite: http.SameSiteLaxMode, // Protege contra requisições forjadas de outros sites (CSRF)
		// Secure: true,               // Descomente esta linha quando colocar em produção com HTTPS
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
