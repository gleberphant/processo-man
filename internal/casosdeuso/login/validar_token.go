package login

import (
	"errors"

	"github.com/gleberphant/ProcessoMan/internal/modelos"
	"github.com/gleberphant/ProcessoMan/internal/repositorios"
)

// verificar se token é valido. retorna error se token não encontraro
func ValidarToken(token modelos.Token) error {

	//verifica se o token é valido
	rows, err := repositorios.Consultar("SELECT uuid, data_criacao FROM tokens WHERE uuid=?; ", token.UUID)

	// se erro na consulta
	if err != nil {
		return err
	}

	defer rows.Close()

	// token inexistente
	if !rows.Next() {
		return errors.New("Token não encontrado")
	}

	// token existente , retorna nulo
	return nil
}
