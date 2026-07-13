package manipuladores

import "net/http"

func (m *ManipuladorAutenticacao) InjetarRotas(mux *http.ServeMux) {

	// AUTENTICACAO
	mux.HandleFunc("GET /login", m.PageLogin)
	mux.HandleFunc("POST /login", m.LoginPost)

}
