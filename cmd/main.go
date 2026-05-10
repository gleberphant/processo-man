package main

import (
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/rotas"
)

func main() {

	router := rotas.Roteador{}

	//configurar servidor
	server := http.Server{
		Addr:    ":8080",
		Handler: router.ConfigurarRotas(),
	}

	fmt.Println("Servidor rodando na porta 8080")
	server.ListenAndServe()

}
