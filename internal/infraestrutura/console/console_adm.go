package console

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gleberphant/ProcessoMan/internal/aplicacao/repositorios"
	"github.com/gleberphant/ProcessoMan/internal/appconfig"
	"github.com/gleberphant/ProcessoMan/internal/dominio/entidades"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
)

type MenuConsole struct {
	Titulo    string
	MapaAcoes map[string]ItemMenu
}

type ItemMenu struct {
	Label string
	Acao  func() *MenuConsole
}

type Console struct {
	config                               *appconfig.AppConfig
	menuPrincipal                        MenuConsole
	submenuBancoDeDados, submenuServicos MenuConsole
}

func NovoConsole(config *appconfig.AppConfig) *Console {
	return &Console{
		config: config,
	}
}

func (c *Console) carregarMenus() {

	c.menuPrincipal = MenuConsole{
		Titulo: "Menu Principal",
		MapaAcoes: map[string]ItemMenu{
			"0": {
				Label: "Sair",
				Acao:  func() *MenuConsole { return nil },
			},
			"1": {
				Label: "Banco de Dados",
				Acao: func() *MenuConsole {
					return &c.submenuBancoDeDados
				},
			},
			"2": {
				Label: "Serviços",
				Acao:  func() *MenuConsole { return &c.submenuBancoDeDados },
			},
		},
	}

	c.submenuBancoDeDados = MenuConsole{
		Titulo: "Menu Banco de Dados",
		MapaAcoes: map[string]ItemMenu{
			"0": {
				Label: "Voltar",
				Acao:  func() *MenuConsole { return &c.menuPrincipal },
			},
			"1": {
				Label: "Rodar Migrações Banco Relacional",
				Acao: func() *MenuConsole {
					ConfigurarBancoRelacional(c.config.ConnDBEntidades)
					return &c.submenuBancoDeDados
				},
			},

			"2": {
				Label: "Configurar Banco Autenticação",
				Acao: func() *MenuConsole {
					ConfigurarBancoChaveValor(c.config.ConnDBAuth)
					return &c.submenuBancoDeDados
				},
			},

			"3": {
				Label: "Inserir Usuario Teste ",
				Acao: func() *MenuConsole {
					InserirUsuarioTeste(c.config.ConnDBEntidades)
					return &c.submenuBancoDeDados
				},
			},

			"4": {
				Label: "Inserir Permissoes Default",
				Acao: func() *MenuConsole {
					InserirPermissoesDefault(c.config.ConnDBAuth)
					return &c.submenuBancoDeDados
				},
			},
		},
	}

	c.submenuServicos = MenuConsole{
		Titulo: "Menu Serviços",
		MapaAcoes: map[string]ItemMenu{
			"0": {
				Label: "Voltar",
				Acao:  func() *MenuConsole { return &c.menuPrincipal },
			},
			"1": {
				Label: "Voltar",
				Acao:  func() *MenuConsole { return &c.menuPrincipal },
			},
		},
	}
}

func (c *Console) Executar() {

	c.carregarMenus()
	fmt.Printf("Iniciando Console de administração do servidor")

	scanner := bufio.NewScanner(os.Stdin)

	menuAtual := &c.menuPrincipal
	for {
		// desenha o menu
		c.ExibeMenu(menuAtual)
		// gerencia entrada
		comando := c.LerComando(scanner)
		// realiza ação
		menuAtual = c.realizaAcaoPrincipal(comando, menuAtual)

		if menuAtual == nil {
			break
		}

	}
}
func (c *Console) ExibeMenu(menu *MenuConsole) {
	fmt.Printf("\n Menu %s", menu.Titulo)
	for key, item := range menu.MapaAcoes {
		fmt.Printf("\n %s - %s ", key, item.Label)
	}

}

func (c *Console) LerComando(scanner *bufio.Scanner) string {
	fmt.Printf("\n Digite a opção :  ")

	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}
	fmt.Printf("Erro ao ler entrada do console")
	return "error"

}

func (c *Console) realizaAcaoPrincipal(comando string, menu *MenuConsole) *MenuConsole {
	// realiza ação
	item, ok := menu.MapaAcoes[comando]

	if !ok {
		fmt.Printf("Opção [%s] inválida ", comando)
		return menu
	}

	proximoMenu := item.Acao()

	return proximoMenu
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

func ConfigurarBancoChaveValor(db *bolt.DB) {

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
