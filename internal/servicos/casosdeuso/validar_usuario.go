package casosdeuso

import (
	"errors"
	"log"

	"github.com/gleberphant/ProcessoMan/internal/modelos"
	"github.com/gleberphant/ProcessoMan/internal/repositorio"
)

// verificar se usuario existe. retorna error se usuario não encontrado
func ValidarUsuario(usuario *modelos.Usuario) error {

	log.Printf("Validando usuario : %s [Senha: %s]", usuario.Email, usuario.Senha)

	rows, err := repositorio.Consultar("SELECT uuid FROM usuarios WHERE email=? and senha=?", usuario.Email, usuario.Senha)

	if err != nil {
		return err
	}

	defer rows.Close()

	// se ha um usuario. então retorna esse usuario
	if !rows.Next() {
		return errors.New("Usuario não encontrado")
	}

	err = rows.Scan(&usuario.UUID)

	// verifica erro na leitura da linha
	if err != nil {
		return err
	}

	log.Printf("Usuario Encontrado: UUDI %s /  EMAIL: %s ", usuario.UUID, usuario.Email)

	return nil

}
