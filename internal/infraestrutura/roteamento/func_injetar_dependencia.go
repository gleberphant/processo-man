package roteamento

import (
	"log"

	"github.com/gleberphant/ProcessoMan/internal/aplicacao/manipuladores"
	"github.com/gleberphant/ProcessoMan/internal/aplicacao/repositorios"
	"github.com/gleberphant/ProcessoMan/internal/dominio/servicos"
)

func (r *Roteador) InjetarDependencias() error {
	// TODO acrescentar tratamento de erros

	// Carregar repositorios com a conexao  de banco  de dados
	log.Printf("Configurando Repositorios")
	repoTokens := repositorios.NovoRepositorioTokenBolt(r.appConfig.ConnDBAuth)
	repoUsuarios := repositorios.NovoRepositorioUsuario(r.appConfig.ConnDBEntidades)
	repoProcessos := repositorios.NovoRepositorioProcesso(r.appConfig.ConnDBEntidades)
	repoTarefas := repositorios.NovoRepositorioTarefa(r.appConfig.ConnDBEntidades)

	// Carregar serviços com os repositorios
	log.Printf("Configurando Servicos")
	servicoAutenticacao := servicos.NovoCDUAutenticacao(repoTokens, repoUsuarios)
	servicoUsuario := servicos.NovoCDUUsuario(repoUsuarios)
	servicoProcesso := servicos.NovoCDUProcesso(repoProcessos, repoTarefas)
	servicoTarefa := servicos.NovoCDUTarefa(repoTarefas, repoProcessos)

	// Carregar  manipuladores HTTP
	log.Printf("Configurando Manipuladores HTTP")
	ManipuladorAutenticacao := manipuladores.NovoManipuladorLogin(servicoAutenticacao)
	ManipuladorUsuario := manipuladores.NovoManipuladorUsuario(servicoUsuario, servicoTarefa)
	ManipuladorProcesso := manipuladores.NovoManipuladorProcesso(servicoProcesso, servicoUsuario)
	ManipuladorTarefa := manipuladores.NovoManipuladorTarefa(servicoTarefa, servicoUsuario)

	// configurar servico autenticacao do roteador
	r.servicoAutenticacao = servicoAutenticacao

	// configurar manipuladores
	r.manipuladores = []IManipulador{
		ManipuladorAutenticacao,
		ManipuladorUsuario,
		ManipuladorProcesso,
		ManipuladorTarefa,
	}

	return nil

}
