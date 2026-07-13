package api

import "net/http"

// ROTAS RESTFUL: RECURSOS ANINHADOS (CLIENTES -> PROCESSOS)

func (m *ApiAreaCliente) InjetarRotasManipuladores(mux *http.ServeMux) {

	mux.HandleFunc("GET /clientes/{cliente_uuid}/processos/{$}", m.AreaClientePageListarMeusProcessos)
	mux.HandleFunc("GET /clientes/{cliente_uuid}/processos/{ProcessoUUID}", m.AreaClientePageVerProcesso)
}
