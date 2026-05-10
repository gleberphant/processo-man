package testes

import (
	"testing"

	"github.com/gleberphant/ProcessoMan/internal/repositorio"
)

func TestBancoDeDados(t *testing.T) {

	//dbManager := infraestrutura.BancoDeDados{}

	conn, err := repositorio.Conectar("../database/sqlite.db")

	if err != nil {
		t.Fatalf("Erro crítico ao conectar no SQLite: %v", err)
	}
	defer conn.Close()

	t.Log("\n LOG :: Conexão com o banco de dados SQLite estabelecida com sucesso!")

	//testar consulta
	//token := "ABC"

	rows, err := repositorio.Consultar("SELECT id, token, permissoes FROM tokens")

	if err != nil {
		t.Fatalf("\n ERRO :: Erro crítico CONSULTAR TABELA: %v", err)
	}

	defer rows.Close()

	tokens := struct {
		id         int
		token      string
		permissoes string
	}{}

	if rows.Next() {
		rows.Scan(&tokens.id, &tokens.token, &tokens.permissoes)

		t.Logf("\n\n [ ID: %d TOKEN: %s PERMISSOES: %s] \n \n", tokens.id, tokens.token, tokens.permissao)

	}

	t.Log("\n LOG :: Consulta realizada com sucesso!")

}
