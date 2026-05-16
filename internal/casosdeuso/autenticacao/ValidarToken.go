package autenticacao

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

// verificar se token é valido. retorna error se token não encontrado
func (a *AutenticacaoCDU) ValidarToken(token entidades.Token) error {

	if token.UUID == uuid.Nil {
		return errors.New("token inválido: UUID não fornecido")
	}

	_, err := a.RepoTokens.BuscarPorUUID(token)

	if err != nil {
		return err
	}

	return nil
}
