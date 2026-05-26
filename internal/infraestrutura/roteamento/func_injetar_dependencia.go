package roteamento

import (
	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
	"github.com/gleberphant/ProcessoMan/internal/dominios/processos"
	"github.com/gleberphant/ProcessoMan/internal/dominios/tarefas"
	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/bancodedados"
)

func (s *Roteador) InjetarDependencias() error {

	// conexao com bancos de dadoss
	db, err := bancodedados.ConectarSQLITE("../database/sqlite.db?_foreign_keys=on")

	if err != nil {
		return err
	}

	dbAuth, err := bancodedados.ConectarSQLITE("../database/autenticacao.db?_foreign_keys=on")

	if err != nil {
		return err
	}

	// cria os repositorios

	tokensRepo := autenticacao.NovoRepositorioToken(dbAuth)

	usuariosRepo := usuarios.NovoRepositorioUsuario(db)
	processoRepo := processos.NovoRepositorioProcesso(db)
	tarefaRepo := tarefas.NovoRepositorioTarefa(db)

	// injeta repositorios nos casos de uso
	cduAutenticacao := autenticacao.NovoCDUAutenticacao(tokensRepo, usuariosRepo)
	cduUsuario := usuarios.NovoCDUUsuario(usuariosRepo)
	cduProcesso := processos.NovoCDUProcesso(processoRepo, tarefaRepo)
	cduTarefa := tarefas.NovoCDUTarefa(tarefaRepo, processoRepo)

	// injeta casos de uso nos manipuladores
	s.LoginManipulador = autenticacao.NovoManipuladorLogin(cduAutenticacao)
	s.ManipuladorUsuario = usuarios.NovoManipuladorUsuario(cduUsuario)
	s.ManipuladorProcesso = processos.NovoManipuladorProcesso(cduProcesso)
	s.ManipuladorTarefa = tarefas.NovoManipuladorTarefa(cduTarefa, cduUsuario)

	return nil

}
