package roteamento

import (
	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
	cduAutenticacao "github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao/casosdeuso"
	"github.com/gleberphant/ProcessoMan/internal/dominios/processos"
	cduProcesso "github.com/gleberphant/ProcessoMan/internal/dominios/processos/casosdeuso"
	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"
	cduUsuario "github.com/gleberphant/ProcessoMan/internal/dominios/usuarios/casosdeuso"
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

	// injeta repositorios nos casos de uso
	cduAutenticacao := cduAutenticacao.NovoCDUAutenticacao(tokensRepo)
	cduUsuario := cduUsuario.NovoCDUUsuario(usuariosRepo)
	cduProcesso := cduProcesso.NovoCDUProcesso(processoRepo)

	// injeta casos de uso nos manipuladores
	s.LoginManipulador = autenticacao.NovoManipuladorLogin(cduAutenticacao)
	s.ManipuladorUsuario = usuarios.NovoManipuladorUsuario(cduUsuario)
	s.ManipuladorProcesso = processos.NovoManipuladorProcesso(cduProcesso)

	return nil

}
