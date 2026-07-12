package intermediarios

import (
	"log"
	"net/http"
)

func LoggerFunc(proximo http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("| Requisao: %s |  Metodo: %s | ", r.URL.Path, r.Method)

		proximo.ServeHTTP(w, r)

		log.Printf(" --------------------------------------------------------- ")

	})
}
