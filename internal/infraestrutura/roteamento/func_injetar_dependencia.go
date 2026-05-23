package roteamento

import (
	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
	"github.com/gleberphant/ProcessoMan/internal/dominios/processos"
	"github.com/gleberphant/ProcessoMan/internal/dominios/tarefas"
	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/bancodedados"
)

func (s *Roteador) InjetarDependencias() error {

	// conexao com o banco de dados
	db, err := bancodedados.ConectarSQLITE()

	if err != nil {
		return err
	}

	// cria os repositorios
	tokensRepo := autenticacao.NovoRepositorioToken(db)
	usuariosRepo := usuarios.NovoRepositorioUsuario(db)
	processoRepo := processos.NovoRepositorioProcesso(db)
	tarefaRepo := tarefas.NovoRepositorioTarefa(db)

	// injeta repositorios nos casos de uso
	cduAutenticacao := autenticacao.NovoCDUAutenticacao(tokensRepo)
	cduUsuario := usuarios.NovoCDUUsuario(usuariosRepo)
	cduProcesso := processos.NovoCDUProcesso(processoRepo, tarefaRepo)
	cduTarefa := tarefas.NovoCDUTarefa(tarefaRepo, processoRepo)

	// injeta casos de uso nos manipuladores
	s.LoginManipulador = autenticacao.NovoManipuladorLogin(cduAutenticacao)
	s.ManipuladorUsuario = usuarios.NovoManipuladorUsuario(cduUsuario)
	s.ManipuladorProcesso = processos.NovoManipuladorProcesso(cduProcesso)
	s.ManipuladorTarefa = tarefas.NovoManipuladorTarefa(cduTarefa)

	return nil

}
