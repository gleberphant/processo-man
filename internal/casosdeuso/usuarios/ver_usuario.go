package usuarios

import (
	"log"

	"github.com/gleberphant/ProcessoMan/internal/modelos"
	"github.com/gleberphant/ProcessoMan/internal/repositorios"
)

func VerCliente(usuario modelos.Usuario) error {

	err := repositorios.Inserir("INSERT INTO usuarios (uuid, nome, email, senha) VALUES (?, ?, ?, ?)",
		usuario.UUID,
		usuario.Nome,
		usuario.Email,
		usuario.Senha,
	)

	if err != nil {
		log.Printf("erro ao criar usuario : %v", err)
		return err
	}

	return nil
}
