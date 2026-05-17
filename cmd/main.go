package main

import (
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/roteamento"
)

func main() {

	r := roteamento.Roteador{}

	r.InjetarDependencias()
	//configurar servidor
	server := http.Server{
		Addr:    ":8080",
		Handler: r.ConfigurarRotas(),
	}

	fmt.Println("Servidor rodando na porta 8080")
	server.ListenAndServe()

}
