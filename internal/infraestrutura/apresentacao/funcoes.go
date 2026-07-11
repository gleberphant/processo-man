package apresentacao

import (
	"net/http"
	"time"
)

// 1. Crie a função de formatação de formatação de data
func usuarioLogado(r *http.Request) string {

	return "usuario logado"
}

func formatarData(data time.Time) string {
	if data.IsZero() {
		return "Sem Data" // ou retorne "" se preferir vazio
	}
	// Formato padrão brasileiro
	return data.Format("02/01/2006")
}

func RedirecionarPaginaAnterior(w http.ResponseWriter, r *http.Request, fallback ...string) {

	destino := "/"

	if len(fallback) > 0 && fallback[0] != "" {
		destino = fallback[0]
	} else if referencia := r.Header.Get("Referer"); referencia != "" {
		destino = referencia
	}

	http.Redirect(w, r, destino, http.StatusSeeOther)
}
