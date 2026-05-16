package autenticacao

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

// gerar token
func (a *AutenticacaoCDU) GerarToken(usuario *entidades.Usuario) (*entidades.Token, error) {

	if usuario == nil {
		return nil, errors.New("usuário inválido")
	}

	repo := a.RepoTokens

	// insere o novo token
	var token = entidades.Token{
		UUID:        uuid.New(),
		UsuarioUUID: usuario.UUID,
		Validade:    "temporario",
	}

	// limpar tokens antigos
	err := repo.Deletar(token)

	if err != nil {
		return nil, err
	}

	//inserir novo token
	err = repo.Criar(token)

	if err != nil {
		return nil, err
	}

	return &token, nil

}
