package usuarios

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios/casosdeuso"
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

// var layout = []string{"../templates/layout/_layout.html", "../templates/layout/_header.html", "../templates/layout/_navbar.html", "../templates/layout/_footer.html"}
type ManipuladorUsuario struct {
	CDUUsario *casosdeuso.Usuario
}

func NovoManipuladorUsuario(casosDeUsoUsuario *casosdeuso.Usuario) *ManipuladorUsuario {

	return &ManipuladorUsuario{
		CDUUsario: casosDeUsoUsuario,
	}
}

func (m *ManipuladorUsuario) PageListar(w http.ResponseWriter, r *http.Request) {

	lista, err := m.CDUUsario.ListarUsuarios()
	if err != nil {
		erroMsg := fmt.Sprintf("Erro :%v", err)
		log.Println(erroMsg)
		http.Error(w, erroMsg, http.StatusInternalServerError)
	}

	apresentacao.ExibirPaginaHTML("pageListarUsuarios.html", w, lista)

}

func (m *ManipuladorUsuario) PageCriar(w http.ResponseWriter, r *http.Request) {

	apresentacao.ExibirPaginaHTML("pageCriarUsuario.html", w, nil)

}

func (m *ManipuladorUsuario) CriarUsuarioPost(w http.ResponseWriter, r *http.Request) {

	UUID, err := uuid.Parse(r.PostFormValue("uuid"))
	var usuario = entidades.Usuario{
		UUID:  UUID,
		Nome:  r.PostFormValue("nome"),
		Email: r.PostFormValue("email"),
		Senha: r.PostFormValue("senha"),
		Perfis: []entidades.Perfil{{
			Nome: r.PostFormValue("perfil"),
		},
		},
	}

	err = m.CDUUsario.CriaUsuario(usuario)

	if err != nil {
		erroMsg := fmt.Sprintf("Erro na criação do usuario:%v", err)
		log.Println(erroMsg)
		http.Error(w, erroMsg, http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/usuario/listar", http.StatusSeeOther)

}

func (m *ManipuladorUsuario) EditarUsuarioPost(w http.ResponseWriter, r *http.Request) {

	var usuario = entidades.Usuario{
		UUID:  uuid.MustParse(r.PostFormValue("uuid")),
		Nome:  r.PostFormValue("nome"),
		Email: r.PostFormValue("email"),
		Senha: r.PostFormValue("senha"),
		Perfis: []entidades.Perfil{{
			Nome: r.PostFormValue("perfil"),
		},
		},
	}
	err := m.CDUUsario.AtualizarUsuario(usuario)

	if err != nil {
		erroMsg := fmt.Sprintf("Erro na criação do usuario:%v", err)
		log.Println(erroMsg)
		http.Error(w, erroMsg, http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/usuario/listar", http.StatusSeeOther)
}

func (m *ManipuladorUsuario) DeletarUsuarioPost(w http.ResponseWriter, r *http.Request) {

	var UUID = r.PostFormValue("uuid")

	err := m.CDUUsario.DeletarUsuario(UUID)

	if err != nil {
		erroMsg := fmt.Sprintf("Erro ao deletar usuario:%v", err)
		log.Println(erroMsg)
		http.Error(w, erroMsg, http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/usuario/listar", http.StatusSeeOther)
}
