package casosdeuso

import (
	"log"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
)

func (u *CDUUsuario) CriaUsuario(usuario entidades.Usuario) error {

	usuario.UUID = uuid.New()

	for i := range usuario.Perfis {
		usuario.Perfis[i].UUID = uuid.New()
	}

	log.Printf("Criando Usuario %v", usuario)
	err := u.RepoUsuarios.Criar(usuario)

	return err

}
