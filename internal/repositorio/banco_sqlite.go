package repositorio

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Driver SQLite
)

type BancoDeDados struct {
	Conexao *sql.DB
}

// conectar ao banco
func Conectar(args ...string) (*sql.DB, error) {

	var caminho string
	// abrir a conexao
	if len(args) > 0 {
		caminho = args[0]
	} else {
		caminho = "../database/sqlite.db"

	}
	var err error

	Conexao, err := sql.Open("sqlite", caminho)

	// 1º verifica se houver erro na conexao
	if err != nil {
		log.Printf("Erro ao conectar-se ao banco de dados : %v", err)
		return nil, err
	}

	// 2º verifica se conexao está ativa
	if err = Conexao.Ping(); err != nil {
		return nil, err
	}

	// retorna a conexao para query

	return Conexao, nil
}

// consultar dado do banco
func Consultar(query string, args ...any) (*sql.Rows, error) {
	//conectar
	conn, err := Conectar()

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// fzer consultar
	res, err := conn.Query(query, args...)

	if err != nil {
		return nil, err
	}
	//retornar resultado

	return res, err
}

// / inserir dados no banco
func Inserir(query string, args ...any) error {
	//conectar
	conn, err := Conectar()

	if err != nil {
		return err
	}

	defer conn.Close()

	// fzer consultar
	_, err = conn.Exec(query, args...)

	if err != nil {
		return err
	}
	//retornar resultado

	return err
}
