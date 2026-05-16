package main

import (
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/rotasHTTP"
)

func main() {

	router := rotasHTTP.Roteador{}

	//configurar servidor
	server := http.Server{
		Addr:    ":8080",
		Handler: router.ConfigurarRotas(),
	}

	fmt.Println("Servidor rodando na porta 8080")
	server.ListenAndServe()

}
