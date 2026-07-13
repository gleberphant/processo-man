package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gleberphant/ProcessoMan/internal/aplicacao/repositorios"
	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/bancodedados"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
)

var URL_BANCO_RELACIONAL string = "../database/sqlite.db"
var URL_BANCO_TOKENS string = "../database/autenticacao.boltdb"

func Configurar() {
	// criar menu de configuração

	fmt.Println("Sistema de Configuração")
	fmt.Println("")
	fmt.Println("..................... MENU ...............")
	fmt.Println(".                                        .")
	fmt.Println(". 1 - Rodar Migrações Banco Relacional   .")
	fmt.Println(". 2 - Configurar Banco Autenticação      .")
	fmt.Println(". 3 - Inserir Usuario Teste    		  .")
	fmt.Println(". 4 - Inserir Permissoes Default         .")
	fmt.Println(".                                        .")
	fmt.Println(". 0 - SAIR                               .")
	fmt.Println("..........................................")
	fmt.Print("Digite a opção: ")

	opcao := 1
	fmt.Println("..........................................")
	switch opcao {
	case 1: // configurar banco relacional com entidades
		dbSqlite := bancodedados.ConectarSQLITE(URL_BANCO_RELACIONAL)
		defer dbSqlite.Close()
		ConfigurarBancoRelacional(dbSqlite)

	case 2: // configurar banco de tokens
		dbBolt := bancodedados.ConectarBBOLT(URL_BANCO_TOKENS)
		defer dbBolt.Close()
		ConfigurarTokensEPermissoesDefault(dbBolt)

	case 3: // inserir usuario de teste
		dbSqlite := bancodedados.ConectarSQLITE()
		defer dbSqlite.Close()
		InserirUsuarioTeste(dbSqlite)

	case 4: // inserir usuario de teste
		dbBolt := bancodedados.ConectarBBOLT(URL_BANCO_TOKENS)
		defer dbBolt.Close()
		ConfigurarTokensEPermissoesDefault(dbBolt)

	default:
	}

}

func ConfigurarBancoRelacional(db *sql.DB) {

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
}

func InserirUsuarioTeste(db *sql.DB) {

	repo := repositorios.NovoRepositorioUsuario(db)

	usuario := entidades.Usuario{
		UUID:  uuid.Nil,
		Nome:  "Nome",
		Email: "teste@teste",
		Senha: "teste",
	}

	repo.Criar(usuario)
}

func ConfigurarTokensEPermissoesDefault(db *bolt.DB) {
	// inserir permissoes default

	permissoes := map[string]map[string]bool{}

	err := db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists([]byte("permissoes"))

		if err != nil {
			return err
		}

		for rota, perfis := range permissoes {
			bytePerfis, err := json.Marshal(perfis)

			if err != nil {
				return err
			}
			rota = strings.ToLower(rota)

			bucket.Put([]byte(rota), bytePerfis)
			log.Printf("chave %s, valor %v", rota, perfis)
		}

		return nil
	})
	if err != nil {
		panic(err)

	}
}

func InserirPermissoesDefault(db *bolt.DB) {
	// inserir permissoes default

	permissoes := map[string]map[string]bool{
		"post:/":              {"admin": true, "colaborador": true},
		"post:/login":         {"admin": true, "colaborador": true},
		"post:/usuarios":      {"admin": true, "colaborador": true},
		"post:/clientes":      {"admin": true, "colaborador": true},
		"post:/colaboradores": {"admin": true, "colaborador": true},
		"post:/processos":     {"admin": true, "colaborador": true, "cliente": true},
		"post:/tarefas":       {"admin": true, "colaborador": true, "cliente": true},
		"get:/":               {"admin": true, "colaborador": true},
		"get:/login":          {"admin": true, "colaborador": true},
		"get:/usuarios":       {"admin": true, "colaborador": true},
		"get:/clientes":       {"admin": true, "colaborador": true},
		"get:/colaboradores":  {"admin": true, "colaborador": true},
		"get:/processos":      {"admin": true, "colaborador": true, "cliente": true},
		"get:/tarefas":        {"admin": true, "colaborador": true, "cliente": true},
	}

	err := db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists([]byte("permissoes"))

		if err != nil {
			return err
		}

		for rota, perfis := range permissoes {
			bytePerfis, err := json.Marshal(perfis)

			if err != nil {
				return err
			}
			rota = strings.ToLower(rota)

			bucket.Put([]byte(rota), bytePerfis)
			log.Printf("chave %s, valor %v", rota, perfis)
		}

		return nil
	})
	if err != nil {
		panic(err)

	}
}
