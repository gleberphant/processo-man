package roteamento

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
	"github.com/gleberphant/ProcessoMan/internal/dominios/processos"
	"github.com/gleberphant/ProcessoMan/internal/dominios/tarefas"
	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"
)

type Roteador struct {
	ManipuladorAutenticacao *autenticacao.ManipuladorLogin
	ManipuladorUsuario      *usuarios.ManipuladorUsuario
	ManipuladorProcesso     *processos.ManipuladorProcesso
	ManipuladorTarefa       *tarefas.ManipuladorTarefa
	Handler                 *http.Handler
}

func NovoRoteador() *Roteador {
	return &Roteador{}
}

func (r *Roteador) Fechar() {
	r.ManipuladorAutenticacao.Fechar()
	r.ManipuladorUsuario.Fechar()
	r.ManipuladorProcesso.Fechar()
	r.ManipuladorTarefa.Fechar()
}
