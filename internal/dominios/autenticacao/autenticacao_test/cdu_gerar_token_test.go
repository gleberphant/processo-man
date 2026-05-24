package autenticacao_test

import (
	"testing"

	"github.com/gleberphant/ProcessoMan/internal/dominios/autenticacao"
	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/bancodedados"
	"github.com/google/uuid"
)

func TestGerarToken(t *testing.T) {

	// conectar banco de dados
	db, _ := bancodedados.ConectarSQLITE()

	// cria os repositorios
	tokensRepo := autenticacao.NovoRepositorioToken(db)

	usuariosRepo := usuarios.NovoRepositorioUsuario(db)

	CDULogin := autenticacao.NovoCDUAutenticacao(tokensRepo, usuariosRepo)

	casosDeTeste := []struct {
		Nome            string // description of this test case
		Entrada         *usuarios.Usuario
		RepostaEsperada *autenticacao.Token
		EsperaFalha     bool
	}{
		{
			Nome:    "Token com usuario valido",
			Entrada: &usuarios.Usuario{UUID: uuid.MustParse("c3f16c59-f802-478b-8c3b-b3b6f20e0af6")},
			//RepostaEsperada: &Token{},
			EsperaFalha: false,
		},
		// {
		// 	NomeDoTeste: "Token com usuario inválido - Deve Falhar",
		// 	Entidade:    &Usuario{UUID: uuid.New()},
		// 	//RepostaEsperada: nil,
		// 	TesteDeFalha: true,
		// },
	}

	for _, teste := range casosDeTeste {
		t.Run(teste.Nome, func(t *testing.T) {

			resposta, err := CDULogin.GerarToken(teste.Entrada)
			// se falhou
			if err != nil {
				// verifica se esperava falha
				if !teste.EsperaFalha {
					t.Errorf("GerarToken() failed: %v", err)
				}
				return
			}
			// teste passou mas esperava falha
			if teste.EsperaFalha {
				t.Errorf("Esperava uma falha na geração do token, mas obteve sucesso")
				return

			}

			// compara resultado
			if resposta == nil {
				//t.Errorf("GerarToken() = %v, want %v", resposta, tt.RepostaEsperada)
				t.Errorf("GerarToken() retornou nil, mas era esperado um ponteiro de token válido")
			}
		})
	}
}
