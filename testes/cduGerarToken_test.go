package testes

import (
	"testing"

	"github.com/gleberphant/ProcessoMan/internal/adaptadores/repositorios"
	"github.com/gleberphant/ProcessoMan/internal/casosdeuso/CasosDeUsoAutenticacao"
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/BancoDeDados"
)

func TestGerarToken(t *testing.T) {

	// conectar banco de dados
	db, _ := BancoDeDados.ConectarSQLITE()

	// cria os repositorios
	tokensRepo := repositorios.NovoRepositorioToken(db)

	usuariosRepo := repositorios.NovoRepositorioUsuario(db)

	CDULogin := CasosDeUsoAutenticacao.NovoCasoDeUsoAutenticacao(tokensRepo, usuariosRepo)

	casosDeTeste := []struct {
		Nome            string // description of this test case
		Entrada         string
		RepostaEsperada *entidades.Token
		EsperaFalha     bool
	}{
		{
			Nome:    "Token com usuario valido",
			Entrada: "c3f16c59-f802-478b-8c3b-b3b6f20e0af6",
			//RepostaEsperada: &entidades.Token{},
			EsperaFalha: false,
		},
		// {
		// 	NomeDoTeste: "Token com usuario inválido - Deve Falhar",
		// 	Entidade:    &entidades.Usuario{UUID: uuid.New()},
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
