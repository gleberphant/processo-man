package area_cliente

import "net/http"

// ROTAS RESTFUL: RECURSOS ANINHADOS (CLIENTES -> PROCESSOS)

func (m *ManipuladorAreaCliente) InjetarRotasProcessos(mux *http.ServeMux) {

	mux.HandleFunc("GET /clientes/{cliente_uuid}/processos/{$}", m.AreaClientePageListarProcessos)
	mux.HandleFunc("GET /clientes/{cliente_uuid}/processos/{ProcessoUUID}", m.AreaClientePageVerProcesso)
}
