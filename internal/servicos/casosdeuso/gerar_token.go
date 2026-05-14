package casosdeuso

import (
	"log"

	"github.com/gleberphant/ProcessoMan/internal/modelos"
	"github.com/gleberphant/ProcessoMan/internal/repositorio"
	"github.com/google/uuid"
)

// gerar token
func GerarToken(usuario modelos.Usuario) (*modelos.Token, error) {

	// Usuario validado então cria um token
	var token = modelos.Token{
		UUID:         uuid.New().String(),
		Usuario_uuid: usuario.UUID,
	}

	err := repositorio.Inserir("INSERT INTO tokens(uuid, usuario_uuid) VALUES(?,?)", token.UUID, token.Usuario_uuid)

	if err != nil {
		log.Printf("Erro na criacao do token: %s", err)
		return nil, err
	}

	return &token, nil

}
