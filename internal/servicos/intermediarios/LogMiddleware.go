package intermediarios

import (
	"log"
	"net/http"
)

func LogMiddleware(proximo http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proximo.ServeHTTP(w, r)

		log.Printf("| Requisao: %s |  Metodo: %s | ",
			r.URL,
			r.Method,
		)
	})
}
