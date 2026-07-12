package intermediarios

import (
	"log"
	"net/http"
)

type Logger struct {
	proximo http.Handler
}

// construtor
func NovoLogger(proximo http.Handler) *Logger {
	return &Logger{
		proximo: proximo,
	}
}

func (l *Logger) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	log.Printf("| Requisao: %s |  Metodo: %s | ",
		req.URL.Path,
		req.Method,
	)

	// chama proximo handler
	l.proximo.ServeHTTP(res, req)
}
