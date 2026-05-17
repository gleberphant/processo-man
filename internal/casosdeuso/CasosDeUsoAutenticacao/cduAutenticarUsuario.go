package CasosDeUsoAutenticacao

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
)

// recebe login e senha e devolver token
func (a *CasosDeUsoAutenticacao) AutenticarUsuario(email string, senha string) (*entidades.Token, error) {

	// Verifica se o usuario existe

	if email == "" {
		return nil, errors.New("Usuario nulo")
	}

	usuarioUUID, err := a.RepoUsuarios.AutenticarUsuario(email, senha)

	if err != nil {
		return nil, err
	}

	// Gera token
	token, err := a.GerarToken(usuarioUUID)
	if err != nil {
		return nil, err
	}

	//retornar o token gerado
	return token, nil
}
