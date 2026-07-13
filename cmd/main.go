package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/apresentacao"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/bancodedados"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/roteamento"
	"github.com/joho/godotenv"
)

func main() {

	log.Printf("Lendo arquivo .env")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Agora a função nativa funciona com o arquivo!
	log.Printf("Carregando variaveis de ambiente")
	porta := os.Getenv("PORTA")

	// Conectar aos bancos de dados
	log.Printf("Conectando bancos de dados autenticador")
	connDBAuth := bancodedados.ConectarBBOLT("../database/autenticacao.boltdb")
	defer connDBAuth.Close()

	log.Printf("Conectando bancos de dados relacional")
	connDBEntidades := bancodedados.ConectarSQLITE("../database/sqlite.db")
	defer connDBEntidades.Close()

	log.Printf("Inciando Roteador")
	roteador := roteamento.NovoRoteador(
		&roteamento.DBConfig{
			ConnDBAuth:      connDBAuth,
			ConnDBEntidades: connDBEntidades,
		},
	)

	log.Printf("INJETAR DEPENDENCIAS")
	roteador.InjetarDependencias()

	log.Printf("INJETAR ROTAS")
	roteador.InjetarRotas()

	log.Printf("INJETAR INTERMEDIARIOS")
	roteador.InjetarIntermediarios()

	log.Printf("Carregar Templates")
	apresentacao.CarregarTemplates()

	server := http.Server{
		Addr:    ":" + porta,
		Handler: roteador.Handler,
	}

	fmt.Println("Servidor rodando na porta 8080")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

	fmt.Println("Encerrando servidor")

}
