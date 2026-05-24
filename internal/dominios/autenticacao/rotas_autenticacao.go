package autenticacao

import "net/http"

func (m *ManipuladorLogin) DefinirRotasAutenticacao() *http.ServeMux {

	mux := http.NewServeMux()

	// AUTENTICACAO
	mux.HandleFunc("GET /login", m.PageLogin)
	mux.HandleFunc("POST /login", m.LoginPost)

	return mux
}
