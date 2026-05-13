package casosdeuso

import (
	"errors"
	"fmt"
	"log"

	"github.com/gleberphant/ProcessoMan/internal/modelos"
	"github.com/gleberphant/ProcessoMan/internal/repositorio"
	"github.com/google/uuid"
)

// gerar token
func Logar(usuario modelos.Usuario) (*modelos.Token, error) {

	// procura se existe usuario
	rows, err := repositorio.Consultar("SELECT uuid, nome FROM usuarios WHERE email=? and senha=?", usuario.Email, usuario.Senha)

	if err != nil {
		return nil, fmt.Errorf("consulta a db: %w", err)
	}

	defer rows.Close()

	// se encontrar o usuario então cria um token
	if rows.Next() {

		var u modelos.Usuario

		err = rows.Scan(&u.UUID, &u.Nome)

		if err != nil {
			return nil, fmt.Errorf("leitura dados usuários: %w", err)
		}

		rows.Close()
		log.Printf("Usuario UUDI %s - NOME %s", u.UUID, u.Nome)

		var token = modelos.Token{
			UUID:  uuid.New(),
			Token: u.UUID,
		}

		_, err := repositorio.Consultar("INSERT INTO tokens(uuid, token) VALUES(?,?)", token.UUID, token.Token)

		if err != nil {
			return nil, fmt.Errorf("criacao do token: %w", err)
		}

		return &token, nil

	}

	return nil, errors.New("Usuario não encontrado")

}
