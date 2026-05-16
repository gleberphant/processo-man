package autenticacao

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
)

// recebe login e senha e devolver token
func (a *AutenticacaoCDU) AutenticarUsuario(usuario *entidades.Usuario) (*entidades.Token, error) {

	// Verifica se o usuario existe

	if usuario == nil {
		return nil, errors.New("Usuario nulo")
	}

	err := a.RepoUsuarios.AutenticarUsuario(usuario)

	if err != nil {
		return nil, err
	}

	// Gera token
	token, err := a.GerarToken(usuario)
	if err != nil {
		return nil, err
	}

	//retornar o token gerado
	return token, nil
}
