package roteamento

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
	"github.com/gleberphant/ProcessoMan/internal/dominios/processos"
	"github.com/gleberphant/ProcessoMan/internal/dominios/tarefas"
	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"
	"github.com/gleberphant/ProcessoMan/internal/manipuladores"
)

type Roteador struct {
	ManipuladorAutenticacao *autenticacao.ManipuladorAutenticacao
	ManipuladorUsuario      *usuarios.ManipuladorUsuario
	ManipuladorProcesso     *processos.ManipuladorProcesso
	ManipuladorTarefa       *tarefas.ManipuladorTarefa

	ManipuladorAreaCliente *manipuladores.ManipuladorAreaCliente
	//IntermediarioAutenticador *intermediarios.Autenticador
	//IntermediarioLogger       *intermediarios.Logger
	Handler *http.Handler
}

func NovoRoteador() *Roteador {

	return &Roteador{
		//	Mux: http.NewServeMux(),
	}
}

func (r *Roteador) Fechar() {
	r.ManipuladorAutenticacao.Fechar()
	r.ManipuladorUsuario.Fechar()
	r.ManipuladorProcesso.Fechar()
	r.ManipuladorTarefa.Fechar()

}
