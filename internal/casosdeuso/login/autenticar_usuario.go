package login

import (
	"log"

	"github.com/gleberphant/ProcessoMan/internal/modelos"
)

func AutenticarUsuario(usuario *modelos.Usuario) (*modelos.Token, error) {

	err := ValidarUsuario(usuario)

	if err != nil {
		log.Printf("Erro ao validar usuario: %v", err)
		return nil, err
	}

	token, err := GerarToken(usuario)

	if err != nil {
		log.Printf("Erro ao gerar token: %v", err)
		return nil, err
	}

	return token, nil

}
