package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gleberphant/ProcessoMan/internal/modelos"
	"github.com/gleberphant/ProcessoMan/internal/repositorio"
	"github.com/gleberphant/ProcessoMan/internal/servicos/casosdeuso"
	"github.com/google/uuid"
)

func main() {

	// carregar banco de dados
	log.Printf("Conectando ao repositório")
	db, err := repositorio.Conectar()

	if err != nil {
		log.Fatalf("Erro na conexao com o banco de dados: %v", err)
	}

	// carregar a lsita de arquivops do diretario de migracoes
	log.Printf("Carregando diretorio de migracões")
	diretorio := "../migracoes/"
	listaDeArquivos, err := os.ReadDir(diretorio)

	if err != nil {
		log.Fatalf("Erro na leitura do diretorios de migracoes : %v", err)
	}

	// para cada arquivo no diretorio migracoes
	log.Printf("Processando arquivos de migracao")
	for _, arquivo := range listaDeArquivos {

		// se a extensão for  diferente de .sql pula o arquivo
		if filepath.Ext(arquivo.Name()) != ".sql" {
			log.Printf("Arquivo ignorado %s", arquivo.Name())
			continue
		}

		// carregar arquivo
		log.Printf("Carregando arquivo %s", diretorio+arquivo.Name())

		migracao, err := os.ReadFile(diretorio + arquivo.Name())

		if err != nil {
			log.Printf("Falha no carregando do arquivo %v", err)
			continue
		}

		// processar arquivo
		log.Printf("Processando arquivo .... ")
		_, err = db.Exec(string(migracao))

		if err != nil {
			log.Printf("Erro na execução : %v", err)
		}

	}

	// inserir USUARIO TESTE
	log.Printf("Inserindo USUARIO TESTE")
	usuario := mockarUsuarios(1)

	// inserir TOKEN TESTE
	log.Printf("Gerar TOKEN TESTE")

	token, err := casosdeuso.GerarToken(usuario)

	if err != nil {
		log.Printf("Erro na execução : %v", err)
	}

	log.Printf("TOKEN GERADO : %s", token.UUID)

}

func mockarUsuarios(args ...int) modelos.Usuario {

	numUsuario := 0

	if len(args) > 0 {
		numUsuario = args[0]
	}

	var usuario modelos.Usuario

	for i := 0; i < numUsuario+1; i++ {

		nome := "teste" + strconv.Itoa(i)

		usuario = modelos.Usuario{
			UUID:  uuid.New().String(),
			Nome:  nome,
			Email: nome + "@teste",
			Senha: nome,
		}
		log.Printf("Inserindo USUARIO %s EMAIL %s", usuario.Nome, usuario.Email)

		err := casosdeuso.InserirUsuario(usuario)

		if err != nil {
			log.Printf("Erro : %v", err)
		}

		log.Printf("Usuario Gerado: %s Email: %s", usuario.UUID, usuario.Email)
	}

	return usuario
}
