package main

import (
	"fmt"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/HTTP"
)

func main() {

	r := HTTP.Roteador{}

	r.InjecaoDependencias()
	//configurar servidor
	server := http.Server{
		Addr:    ":8080",
		Handler: r.ConfigurarRotas(),
	}

	fmt.Println("Servidor rodando na porta 8080")
	server.ListenAndServe()

}
