package casosdeuso

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

// gerar token
func (a *CasosDeUsoAutenticacao) GerarToken(usuarioUUID string) (*entidades.Token, error) {

	if usuarioUUID == "" {
		return nil, errors.New("usuário inválido")
	}

	// limpar tokens antigos no repositorio
	err := a.RepoTokens.DeletarPorUsuarioUUID(usuarioUUID)

	if err != nil {
		return nil, err
	}

	// inserir o novo token no repositorio
	token, err := a.RepoTokens.Criar(entidades.Token{
		UUID:        uuid.New(),
		UsuarioUUID: uuid.MustParse(usuarioUUID),
		Validade:    "temporario",
	})

	if err != nil {
		return nil, err
	}

	return token, nil

}
