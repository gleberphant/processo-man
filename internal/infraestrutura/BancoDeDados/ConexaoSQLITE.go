package BancoDeDados

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Driver SQLite
)

// conectar ao banco
func ConectarSQLITE(args ...string) (*sql.DB, error) {

	var caminho string
	// abrir a conexao
	if len(args) > 0 {
		caminho = args[0]
	} else {
		caminho = "../database/sqlite.db?_foreign_keys=on"

	}
	var err error

	conn, err := sql.Open("sqlite", caminho)

	// 1º verifica se houver erro na conexao
	if err != nil {
		log.Printf("Erro ao conectar-se ao banco de dados : %v", err)
		return nil, err
	}

	// 2º verifica se conexao está ativa
	if err = conn.Ping(); err != nil {
		return nil, err
	}

	// retorna a conexao para query

	return conn, nil
}
