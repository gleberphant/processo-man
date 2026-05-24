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

type ViewModelUsuario struct {
	UUID     string
	Usuarios interface{}
	Tarefas  interface{}
	Anexos   interface{}
}

func NovoManipuladorUsuario(casosDeUsoUsuario *CDUUsuario) *ManipuladorUsuario {

	return &ManipuladorUsuario{
		cduUsario: casosDeUsoUsuario,
	}
}

func (m *ManipuladorUsuario) PageCriarUsuario(w http.ResponseWriter, r *http.Request) {

	viewModel := ViewModelUsuario{}

	apresentacao.ExibirPaginaHTML("usuario/page-criar-usuario.html", w, viewModel)

}

func (m *ManipuladorUsuario) PageListarUsuarios(w http.ResponseWriter, r *http.Request) {

	lista, err := m.cduUsario.ListarUsuarios()
	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro PageListar:%v", err))
		return
	}

	viewModel := ViewModelUsuario{
		Usuarios: lista,
	}

	apresentacao.ExibirPaginaHTML("usuario/page-listar-usuario.html", w, viewModel)

}

func (m *ManipuladorUsuario) PageEditarUsuario(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	usuario, err := m.cduUsario.BuscarUsuarioPorUUID(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro PageEditar:%v", err))
	}

	viewModel := ViewModelUsuario{
		UUID:     strUUID,
		Usuarios: usuario,
	}

	apresentacao.ExibirPaginaHTML("usuario/page-criar-usuario.html", w, viewModel)

}

func (m *ManipuladorUsuario) PageVerUsuario(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PathValue("UUID")

	usuario, err := m.cduUsario.BuscarUsuarioPorUUID(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Page Editar:%v", err))
	}

	viewModel := ViewModelUsuario{
		Usuarios: usuario,
		Anexos:   []string{"documento1", "document2"},
		//Tarefas:  []tarefas.Tarefa{},
	}

	apresentacao.ExibirPaginaHTML("usuario/page-ver-usuario.html", w, viewModel)

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

	http.Redirect(w, r, "/usuarios", http.StatusSeeOther)

}

func (m *ManipuladorUsuario) EditarUsuarioPost(w http.ResponseWriter, r *http.Request) {

	strUUID := r.PostFormValue("uuid") //r.PathValue("UUID")

	UUID, err := uuid.Parse(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro editar Processo:%v", err))
	}

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

	http.Redirect(w, r, "/usuarios", http.StatusSeeOther)
}

func (m *ManipuladorUsuario) DeletarUsuarioPost(w http.ResponseWriter, r *http.Request) {

	var strUUID = r.PathValue("UUID") //r.PostFormValue("uuid")

	err := m.cduUsario.DeletarUsuario(strUUID)

	if err != nil {
		apresentacao.ExibirErro(w, fmt.Sprintf("Erro Editar Usuario:%v", err))
	}

	http.Redirect(w, r, "/usuarios", http.StatusSeeOther)
}
