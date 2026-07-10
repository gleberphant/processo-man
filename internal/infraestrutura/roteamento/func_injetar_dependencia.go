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
	tokensRepo := autenticacao.NovoRepositorioTokenBolt(connDBAuth)
	usuariosRepo := usuarios.NovoRepositorioUsuario(connDBEntidades)
	processoRepo := processos.NovoRepositorioProcesso(connDBEntidades)
	tarefaRepo := tarefas.NovoRepositorioTarefa(connDBEntidades)

	// injeta repositorios nos casos de uso
	log.Printf("Configurando servicos")
	cduAutenticacao := autenticacao.NovoCDUAutenticacao(tokensRepo, usuariosRepo)
	cduUsuario := usuarios.NovoCDUUsuario(usuariosRepo)
	cduProcesso := processos.NovoCDUProcesso(processoRepo, tarefaRepo)
	cduTarefa := tarefas.NovoCDUTarefa(tarefaRepo, processoRepo)

	// injeta casos de uso nos manipuladores
	log.Printf("Configurando Manipuladores HTTP")
	r.ManipuladorAutenticacao = autenticacao.NovoManipuladorLogin(cduAutenticacao)
	r.ManipuladorUsuario = usuarios.NovoManipuladorUsuario(cduUsuario, cduTarefa)
	r.ManipuladorProcesso = processos.NovoManipuladorProcesso(cduProcesso, cduUsuario)
	r.ManipuladorTarefa = tarefas.NovoManipuladorTarefa(cduTarefa, cduUsuario)

	return nil

}
