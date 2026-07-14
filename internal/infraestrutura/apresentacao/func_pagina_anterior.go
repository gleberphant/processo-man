package apresentacao

import (
	"net/http"
)

func RedirecionarPaginaAnterior(w http.ResponseWriter, r *http.Request, fallback ...string) {

	destino := "/"

	if len(fallback) > 0 && fallback[0] != "" {
		destino = fallback[0]
	} else if referencia := r.Header.Get("Referer"); referencia != "" {
		destino = referencia
	}

	http.Redirect(w, r, destino, http.StatusSeeOther)
}
