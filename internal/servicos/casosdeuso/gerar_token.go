package casosdeuso

import (
	"log"

	"github.com/gleberphant/ProcessoMan/internal/modelos"
	"github.com/gleberphant/ProcessoMan/internal/repositorio"
	"github.com/google/uuid"
)

// gerar token
func GerarToken(usuario modelos.Usuario) (*modelos.Token, error) {

	// primeiro limpa tokens antigos do mesmo usuário para não gerar insegurança
	err := repositorio.Inserir("DELETE FROM tokens WHERE usuario_uuid=?", usuario.UUID)

	if err != nil {
		log.Printf("Erro na LIMPEZA de tokens antigos : %s", err)
		return nil, err
	}

	// insere o novo token
	var token = modelos.Token{
		UUID:         uuid.New().String(),
		Usuario_uuid: usuario.UUID,
		Validade:     "temporario",
	}

	err = repositorio.Inserir("INSERT INTO tokens(uuid, usuario_uuid, validade) VALUES(?,?, ?)", token.UUID, token.Usuario_uuid, token.Validade)

	if err != nil {
		log.Printf("Erro na criacao do token: %s", err)
		return nil, err
	}

	return &token, nil

}
