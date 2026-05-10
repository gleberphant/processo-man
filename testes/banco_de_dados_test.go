package testes

import (
	"testing"

	"github.com/gleberphant/ProcessoMan/internal/repositorio"
)

func TestBancoDeDados(t *testing.T) {

	//dbManager := infraestrutura.BancoDeDados{}

	db, err := repositorio.Conectar()

	if err != nil {
		t.Fatalf("Erro crítico ao conectar no SQLite: %v", err)
	}

	// Se chegou aqui, o Ping() passou e o arquivo ./sistema.db foi criado/aberto
	t.Log("\nConexão com o banco de dados SQLite estabelecida com sucesso!")

	// Fechar a conexão ao encerrar o programa
	defer db.Close()

}
