package casosdeuso

import (
	"log"

	"github.com/gleberphant/ProcessoMan/internal/modelos"
	"github.com/gleberphant/ProcessoMan/internal/repositorio"
)

func InserirUsuario(usuario modelos.Usuario) error {

	err := repositorio.Inserir("INSERT INTO usuarios (uuid, nome, email, senha) VALUES (?, ?, ?, ?)",
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
