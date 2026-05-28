package autenticacao

import "net/http"

func (m *ManipuladorLogin) DefinirRotasAutenticacao(mux *http.ServeMux) {

	// AUTENTICACAO
	mux.HandleFunc("GET /login", m.PageLogin)
	mux.HandleFunc("POST /login", m.LoginPost)

}
