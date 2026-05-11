package casosdeuso

import (
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/repositorio"
	"github.com/google/uuid"
)

// retorna o token gerado
func Logar(usuario entidades.Usuario) (uuid.UUID, error) {

	// chamar repositorio e pesquisar se existe usuarios

	// procura se existe usuario
	query, err := repositorio.Consultar("SELECT uuid FROM usuario WHERE email=? and senha=?", usuario.Email, usuario.Senha)

	if err != nil {
		return uuid.UUID{}, err
	}

	// se retornou o usuario então cria um token

	if query.Next() {
		token, _ := uuid.NewUUID()

		_, err := repositorio.Consultar("INSERT INTO tokens(uuid, token) VALUES(?,?)", token, token)

		if err != nil {
			return uuid.UUID{}, err
		}

		return token, nil

	}

}
