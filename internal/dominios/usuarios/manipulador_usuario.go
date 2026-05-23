package usuarios

import (
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/google/uuid"
)

type ManipuladorUsuario struct {
	cduUsario *CDUUsuario
}

func NovoManipuladorUsuario(casosDeUsoUsuario *CDUUsuario) *ManipuladorUsuario {

	return &ManipuladorUsuario{
		cduUsario: casosDeUsoUsuario,
	}
}

func (m *ManipuladorUsuario) PageListar(w http.ResponseWriter, r *http.Request) {

	lista, err := m.cduUsario.ListarUsuarios()
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro PageListar:%v", err))
		return
	}

	apresentacao.ExibirPaginaHTML("usuario/page-listar-usuario.html", w, lista)

}

func (m *ManipuladorUsuario) PageCriar(w http.ResponseWriter, r *http.Request) {

	apresentacao.ExibirPaginaHTML("usuario/page-criar-usuario.html", w, Usuario{})

}

func (m *ManipuladorUsuario) PageEditar(w http.ResponseWriter, r *http.Request) {

	uuidStr := r.URL.Query().Get("uuid")

	usuario, err := m.cduUsario.BuscarUsuarioPorUUID(uuidStr)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro PageEditar:%v", err))
	}

	apresentacao.ExibirPaginaHTML("usuario/page-criar-usuario.html", w, usuario)

}

func (m *ManipuladorUsuario) CriarUsuarioPost(w http.ResponseWriter, r *http.Request) {

	var usuario = Usuario{
		Nome:  r.PostFormValue("nome"),
		Email: r.PostFormValue("email"),
		Senha: r.PostFormValue("senha"),
		Perfis: []Perfil{{
			Nome: r.PostFormValue("perfil"),
		},
		},
	}

	err := m.cduUsario.CriaUsuario(usuario)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Criar Usuario:%v", err))
	}

	http.Redirect(w, r, "/usuario/listar", http.StatusSeeOther)

}

func (m *ManipuladorUsuario) EditarUsuarioPost(w http.ResponseWriter, r *http.Request) {

	UUID, err := uuid.Parse(r.PostFormValue("uuid"))

	var usuario = Usuario{
		UUID:  UUID,
		Nome:  r.PostFormValue("nome"),
		Email: r.PostFormValue("email"),
		Senha: r.PostFormValue("senha"),
		Perfis: []Perfil{{
			Nome: r.PostFormValue("perfil"),
		},
		},
	}
	err = m.cduUsario.EditarUsuario(usuario)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Usuario:%v", err))
	}

	http.Redirect(w, r, "/usuario/listar", http.StatusSeeOther)
}

func (m *ManipuladorUsuario) DeletarUsuarioPost(w http.ResponseWriter, r *http.Request) {

	var UUID = r.PostFormValue("uuid")

	err := m.cduUsario.DeletarUsuario(UUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Usuario:%v", err))
	}

	http.Redirect(w, r, "/usuario/listar", http.StatusSeeOther)
}
