package entidades

import "github.com/google/uuid"

type Perfil struct {
	ID              uuid.UUID   `json:"id,omitempty"  db:"id"`
	Nome            string      `json:"nome,omitempty"  db:"nome"`
	ListaPermissoes []Permissao `json:"lista_permissoes,omitempty"  db:"-"`
}
