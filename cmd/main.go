package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
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

	log.Printf("Carregar Templates")
	apresentacao.CarregarTemplates()

	server := http.Server{
		Addr:    ":8080",
		Handler: *r.Handler,
	}

	fmt.Println("Servidor rodando na porta 8080")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

	fmt.Println("Encerrando servidor")

}
