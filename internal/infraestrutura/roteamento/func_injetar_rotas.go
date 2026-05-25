package roteamento

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
	"github.com/gleberphant/ProcessoMan/internal/dominios/paginasestaticas"
)

// configurar as rotas e devolver MUX configurado
func (s *Roteador) ConfigurarRotas() http.Handler {

	// configurar mux
	mux := http.NewServeMux()

	// PAGINAS ESTÁTICAS
	mux.HandleFunc("/{$}", paginasestaticas.Index)
	mux.HandleFunc("/", paginasestaticas.Pagina404)

	// SUB-ROTEADORES
	autenticacaoMux := s.LoginManipulador.DefinirRotasAutenticacao()
	usuariosMux := s.ManipuladorUsuario.DefinirRotasUsuarios()
	clientesMux := s.ManipuladorUsuario.DefinirRotasClientes()
	colaboradoresMux := s.ManipuladorUsuario.DefinirRotasColaboradores()
	processosMux := s.ManipuladorProcesso.DefinirRotasProcessos()
	tarefasMux := s.ManipuladorTarefa.DefinirRotasTarefas()

	// INJETANDO SUB-ROTEADORES NO ROTEADOR PRINCIPAL (ANINHAMENTO)
	mux.Handle("/login", autenticacaoMux)
	mux.Handle("/usuarios/clientes/", http.StripPrefix("/usuarios/clientes", clientesMux))
	mux.Handle("/usuarios/colaboradores/", http.StripPrefix("/usuarios/colaboradores", colaboradoresMux))
	mux.Handle("/usuarios/", http.StripPrefix("/usuarios", usuariosMux))
	mux.Handle("/processos/", http.StripPrefix("/processos", processosMux))
	mux.Handle("/tarefas/", http.StripPrefix("/tarefas", tarefasMux))

	// INJETA INTERMEDIÁRIOS - Middlewares
	roteador := autenticacao.LogIntermediario(autenticacao.AutenticadorIntermediario(mux, s.LoginManipulador.CDUAutenticacao))
	return roteador

}
