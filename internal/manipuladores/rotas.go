package manipuladores

import "net/http"

// ROTAS RESTFUL: RECURSOS ANINHADOS (CLIENTES -> PROCESSOS)

func (m *ManipuladorAreaCliente) InjetarRotasManipuladores(mux *http.ServeMux) {

	mux.HandleFunc("GET /clientes/{cliente_uuid}/processos/{$}", m.PageListarMeusProcessos)
	mux.HandleFunc("GET /clientes/{cliente_uuid}/processos/{ProcessoUUID}", m.AreaClientePageVerProcesso)
}
