package controladores

import (
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/gleberphant/ProcessoMan/internal/modelos"
	"github.com/gleberphant/ProcessoMan/internal/servicos/casosdeuso"
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
		log.Printf("Erro ao carregar template: %v", err.Error())
		http.Error(w, "Erro ao carregar template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// executar template
	err = tmpl.Execute(w, dados)
	if err != nil {
		log.Printf("erro ao executar template")
		http.Error(w, "Erro ao renderizar pagina", http.StatusInternalServerError)
		return
	}

}

// funcao para logar
func LoginPost(w http.ResponseWriter, r *http.Request) {

	var usuario = modelos.Usuario{
		Email: r.PostFormValue("email"),
		Senha: r.PostFormValue("senha"),
	}

	err := casosdeuso.ValidarUsuario(&usuario)

	if err != nil {
		log.Printf("Erro ao validar usuario: %v", err)
		http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso Negado. Informe E-mail e Senha Válidos."), http.StatusSeeOther)

		return
	}

	token, err := casosdeuso.GerarToken(usuario)

	if err != nil {
		log.Printf("Erro ao gerar token de acesso : %v", err)
		http.Redirect(w, r, "/login?msg="+url.QueryEscape("Acesso Negado. Erro ao gerar token de acesso tente novamente."), http.StatusSeeOther)
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

	// redireciona para o index
	log.Printf("Token Gerado: [%s] ", token.UUID)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
