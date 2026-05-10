package testes

import (
	"net/http"
	"testing"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/rotas"
)

func TestServidorHTTP(t *testing.T) {

	router := rotas.Roteador{}

	//configurar servidor
	server := http.Server{
		Addr:    ":8080",
		Handler: router.ConfigurarRotas(),
	}

	t.Log("Servidor rodando na porta 8080")
	server.ListenAndServe()

}
