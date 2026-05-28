package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/roteamento"
)

func main() {

	log.Printf("Configurar roteador")
	r := roteamento.Roteador{}

	defer r.Fechar()
	log.Printf("INJETAR DEPENDENCIAS")
	r.InjetarDependencias()
	log.Printf("INJETAR ROTAS")
	r.InjetarRotas()

	server := http.Server{
		Addr:    ":8080",
		Handler: *r.Handler,
	}

	fmt.Println("Servidor rodando na porta 8080")

	server.ListenAndServe()

}
