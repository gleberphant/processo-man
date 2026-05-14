package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gleberphant/ProcessoMan/internal/repositorio"
)

func main() {

	// carregar banco de dados
	log.Printf("Conectando ao repositório")
	db, err := repositorio.Conectar()

	if err != nil {
		log.Fatalf("Erro na conexao com o banco de dados: %v", err)
	}

	// carregar a lsita de arquivops do diretario de migracoes
	log.Printf("Carregando lista de arquivos de migracao")
	diretorio := "migracoes"
	listaDeArquivos, err := os.ReadDir(diretorio)

	if err != nil {
		log.Fatalf("Erro nao ler diretorios de migracoes : %v", err)
	}

	// para cada arquivo no diretorio migracoes
	log.Printf("Processando arquivos de migracao")
	for _, arquivo := range listaDeArquivos {

		// se a extensão for  diferente de .sql pula o arquivo
		if filepath.Ext(arquivo.Name()) != ".sql" {
			log.Printf("Arquivo ignorado %s", arquivo.Name())
			continue
		}

		//carregar arquivo
		log.Printf("Carregando arquivo %s", diretorio+arquivo.Name())

		migracao, err := os.ReadFile(diretorio + arquivo.Name())

		if err != nil {
			log.Printf("Falha no carregando do arquivo %v", err)
			continue
		}

		//processar arquivo
		log.Printf("Processando arquivo %s", migracao)
		_, err = db.Exec(string(migracao))

		if err != nil {

			log.Printf("Erro na execução : %v", err)
		}
	}

}
