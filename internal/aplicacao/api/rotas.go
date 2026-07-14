package api

import "net/http"

// ROTAS RESTFUL: RECURSOS ANINHADOS (CLIENTES -> PROCESSOS)

func (m *ApiAreaCliente) InjetarRotasManipuladores(mux *http.ServeMux) {

	//mux.HandleFunc("GET /clientes/{cliente_uuid}/processos/{$}", m.AreaClientePageListarMeusProcessos)
	//mux.HandleFunc("GET /clientes/{cliente_uuid}/processos/{ProcessoUUID}", m.AreaClientePageVerProcesso)

	// rotas da area do colaborador -- temporario depois mudar para o local correto
	// mux.HandleFunc("GET /colaboradores/{$}", m.PageListarColaboradores)
	// mux.HandleFunc("GET /colaboradores/criar", m.PageCriarColaborador)
	//mux.HandleFunc("POST /colaboradores/criar", m.CriarColaboradorPost)
}
