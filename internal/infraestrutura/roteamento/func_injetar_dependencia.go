package roteamento

import (
	"log"

	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
	"github.com/gleberphant/ProcessoMan/internal/dominios/processos"
	"github.com/gleberphant/ProcessoMan/internal/dominios/tarefas"
	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"
	"go.etcd.io/bbolt"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/bancodedados"
)

func (s *Roteador) InjetarDependencias() error {

	// cria os repositorios
	log.Printf("Carregando bando de autenticacao tokens bolt")

	dbAuth, err := bbolt.Open("../database/autenticacao.boltdb", 0600, nil)

	if err != nil {
		return err
	}

	tokensRepo := autenticacao.NovoRepositorioTokenBolt(dbAuth)

	// conexao com bancos de dadoss
	log.Printf("Carregando bando relacional")
	db, err := bancodedados.ConectarSQLITE("../database/sqlite.db?_foreign_keys=on")

	if err != nil {
		return err
	}

	log.Printf("repositorio tokens usuarios")
	usuariosRepo := usuarios.NovoRepositorioUsuario(db)
	log.Printf("repositorio tokens processos")
	processoRepo := processos.NovoRepositorioProcesso(db)
	log.Printf("repositorio tokens tarefas")
	tarefaRepo := tarefas.NovoRepositorioTarefa(db)

	// injeta repositorios nos casos de uso
	cduAutenticacao := autenticacao.NovoCDUAutenticacao(tokensRepo, usuariosRepo)
	cduUsuario := usuarios.NovoCDUUsuario(usuariosRepo)
	cduProcesso := processos.NovoCDUProcesso(processoRepo, tarefaRepo)
	cduTarefa := tarefas.NovoCDUTarefa(tarefaRepo, processoRepo)

	// injeta casos de uso nos manipuladores
	s.ManipuladorAutenticacao = autenticacao.NovoManipuladorLogin(cduAutenticacao)
	s.ManipuladorUsuario = usuarios.NovoManipuladorUsuario(cduUsuario, cduTarefa)
	s.ManipuladorProcesso = processos.NovoManipuladorProcesso(cduProcesso, cduUsuario)
	s.ManipuladorTarefa = tarefas.NovoManipuladorTarefa(cduTarefa, cduUsuario)

	return nil

}
