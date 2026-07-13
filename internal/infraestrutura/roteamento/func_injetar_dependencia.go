package roteamento

import (
	"log"

	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
	"github.com/gleberphant/ProcessoMan/internal/dominios/processos"
	"github.com/gleberphant/ProcessoMan/internal/dominios/tarefas"
	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/bancodedados"
)

func (r *Roteador) InjetarDependencias() error {

	// Conectar aos bancos de dados

	log.Printf("Conectando bancos de dados ")
	connDBAuth := bancodedados.ConectarBBOLT("../database/autenticacao.boltdb")
	connDBEntidades := bancodedados.ConectarSQLITE("../database/sqlite.db")

	// injeta banco de dados nos repositórios
	log.Printf("Configurando repositorios")
	repoTokens := autenticacao.NovoRepositorioTokenBolt(connDBAuth)
	repoUsuarios := usuarios.NovoRepositorioUsuario(connDBEntidades)
	repoProcessos := processos.NovoRepositorioProcesso(connDBEntidades)
	repoTarefas := tarefas.NovoRepositorioTarefa(connDBEntidades)

	// injeta repositorios nos casos de uso
	log.Printf("Configurando Servicos")
	servicoAutenticacao := autenticacao.NovoCDUAutenticacao(repoTokens, repoUsuarios)
	servicoUsuario := usuarios.NovoCDUUsuario(repoUsuarios)
	servicoProcesso := processos.NovoCDUProcesso(repoProcessos, repoTarefas)
	servicoTarefa := tarefas.NovoCDUTarefa(repoTarefas, repoProcessos)

	// injeta casos de uso nos manipuladores
	log.Printf("Configurando Manipuladores HTTP")
	r.ManipuladorAutenticacao = autenticacao.NovoManipuladorLogin(servicoAutenticacao)
	r.ManipuladorUsuario = usuarios.NovoManipuladorUsuario(servicoUsuario, servicoTarefa)
	r.ManipuladorProcesso = processos.NovoManipuladorProcesso(servicoProcesso, servicoUsuario)
	r.ManipuladorTarefa = tarefas.NovoManipuladorTarefa(servicoTarefa, servicoUsuario)

	// manipuladores de API

	// injetar intermediarios
	//	r.IntermediarioAutenticador = intermediarios.NovoAutenticador(*r.Handler, cduAutenticacao)
	//	r.IntermediarioLogger = intermediarios.NovoIntermediarioLogger(r.IntermediarioAutenticador)

	return nil

}
