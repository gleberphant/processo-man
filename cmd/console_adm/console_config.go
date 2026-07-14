package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gleberphant/ProcessoMan/internal/appconfig"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/bancodedados"
	"github.com/gleberphant/ProcessoMan/internal/infraestrutura/console"
	"github.com/joho/godotenv"
)

func main() {

	var appConfig *appconfig.AppConfig
	var ok bool
	var URLBancoChaveValor, URLBancoRelacional, DiretorioRaiz, PortaServidor string

	// CONFIGURAR APLICAÇÃO
	// definindo HOME
	DiretorioRaiz, ok = os.LookupEnv("HOME")
	if !ok {
		DiretorioRaiz = appconfig.GetDiretorioRaiz()
		os.Setenv("HOME", DiretorioRaiz)
	}

	// Carregar arquivo dot .env
	log.Printf("Lendo arquivo .env")
	err := godotenv.Load(filepath.Join(DiretorioRaiz, ".env"))
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// CARREGAR URL DOS BANCOS DE DADOS
	if os.Getenv("AMBIENTE") == "LOCAL" {
		URLBancoChaveValor = filepath.Join(DiretorioRaiz, os.Getenv("URLBancoChaveValor"))
		URLBancoRelacional = filepath.Join(DiretorioRaiz, os.Getenv("URLBancoRelacional"))
	} else {
		URLBancoChaveValor = os.Getenv("URLBancoChaveValor")
		URLBancoRelacional = os.Getenv("URLBancoRelacional")
	}

	// CARREGAR PORTA
	log.Printf("Carregando variaveis de ambiente")
	PortaServidor = os.Getenv("PORTA")

	// Conectar aos bancos de dados
	log.Printf("Conectando bancos de dados autenticador")
	connDBAuth := bancodedados.ConectarBBOLT(URLBancoChaveValor)
	defer connDBAuth.Close()

	log.Printf("Conectando bancos de dados relacional")
	connDBEntidades := bancodedados.ConectarSQLITE(URLBancoRelacional)
	defer connDBEntidades.Close()

	// configuracao
	log.Printf("Configurando aplicação")
	appConfig = appconfig.NovoAppConfig(
		DiretorioRaiz,
		PortaServidor,
		connDBAuth,
		connDBEntidades,
	)

	log.Printf("Configura console dinamico")
	consoleAdm := console.NovoConsole(appConfig)

	consoleAdm.Executar()
}
