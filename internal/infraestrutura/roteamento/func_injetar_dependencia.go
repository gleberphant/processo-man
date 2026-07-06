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
	dbAuth, err := bancodedados.ConectarBBOLT("../database/autenticacao.boltdb")
	if err != nil {
		return err
	}

	db, err := bancodedados.ConectarSQLITE("../database/sqlite.db")
	if err != nil {
		return err
	}

	// injeta banco de dados nos repositórios
	log.Printf("Configurando repositorios")
	tokensRepo := autenticacao.NovoRepositorioTokenBolt(dbAuth)
	usuariosRepo := usuarios.NovoRepositorioUsuario(db)
	processoRepo := processos.NovoRepositorioProcesso(db)
	tarefaRepo := tarefas.NovoRepositorioTarefa(db)

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
