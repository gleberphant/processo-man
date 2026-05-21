package roteamento

import (
	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
	"github.com/gleberphant/ProcessoMan/internal/dominios/processos"
	"github.com/gleberphant/ProcessoMan/internal/dominios/tarefas"
	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"
)

type Roteador struct {
	LoginManipulador    *autenticacao.ManipuladorLogin
	ManipuladorUsuario  *usuarios.ManipuladorUsuario
	ManipuladorProcesso *processos.ManipuladorProcesso
	ManipuladorTarefa   *tarefas.ManipuladorTarefa
}
