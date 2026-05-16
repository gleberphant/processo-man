package repositorios

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/infraestrutura"
	"github.com/gleberphant/ProcessoMan/internal/modelos"
)

type TokenSqlite struct {
}

// consultar dado do banco
func (t *TokenSqlite) Consultar(token modelos.Token) (*modelos.Token, error) {
	rows, err := infraestrutura.Consultar("SELECT usuario_id, data_criacao FROM tokens WHERE uuid=?; ", token.UUID)

	// se erro na consulta
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("Token não encontrado")
	}

	rows.Scan(&token.UsuarioUUID, &token.DataCriacao)
	return &token, nil

}

// / inserir dados no banco
func (t *TokenSqlite) Inserir(query string, args ...any) error {
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
