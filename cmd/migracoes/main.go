package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gleberphant/ProcessoMan/internal/dominios/usuarios"
	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
	// Assumindo que InserirUsuario esteja aqui
)

func main() {

	// Configurar Banco de Dados SQLITE
	// dbSqlite, err := bancodedados.ConectarSQLITE()
	// if err != nil {
	// 	log.Fatalf("Erro na conexao com SQLITE: %v", err)
	// }
	// defer dbSqlite.Close()

	// // rodar migracoes
	// RodarMigrations(dbSqlite)

	// configurar permissões
	dbBolt, err := bolt.Open("../database/autenticacao.boltdb", 0600, nil)

	if err != nil {
		log.Fatalf("Erro na conexao com BoltDB: %v", err)

	}

	defer dbBolt.Close()

	ConfigurarTokensEPermissoesDefault(dbBolt)

}

func RodarMigrations(db *sql.DB) {

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

	repo := usuarios.NovoRepositorioUsuario(db)

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
