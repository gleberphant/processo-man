package casosdeuso_test

import (
	"testing"

	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"
	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios/casosdeuso"
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/bancodedados"
	"github.com/google/uuid"
)

func TestCriaUsuario(t *testing.T) {

	db, err := bancodedados.ConectarSQLITE("../database/sqlite.db")
	if err != nil {
		t.Fatalf("Erro conexao banco de dados %v", err)
	}

	casosDeUsoUsuarios := casosdeuso.NovoCDUUsuario(
		usuarios.NovoRepositorioUsuario(db),
	)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		usuario entidades.Usuario
		wantErr bool
	}{
		{name: "Criar um usuario comum",
			usuario: entidades.Usuario{
				UUID:  uuid.New(),
				Nome:  "Novo Usuario",
				Email: "novo@novo",
				Senha: "novo@novo",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := casosDeUsoUsuarios.CriaUsuario(tt.usuario)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CriaUsuario() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CriaUsuario() succeeded unexpectedly")
			}
		})
	}
}
