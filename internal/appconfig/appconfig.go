package appconfig

import (
	"database/sql"

	"go.etcd.io/bbolt"
)

// struct que carrega as configurações GLOBAIS da aplicação
type AppConfig struct {
	DiretorioRaiz   string
	PortaServidor   string
	ConnDBAuth      *bbolt.DB // conexao chave valor
	ConnDBEntidades *sql.DB   // conexao relacional
}

func NovoAppConfig(DiretorioRaiz string, PortaServidor string, ConnDBAuth *bbolt.DB, ConnDBEntidades *sql.DB) *AppConfig {
	return &AppConfig{
		DiretorioRaiz:   DiretorioRaiz,
		PortaServidor:   PortaServidor,
		ConnDBAuth:      ConnDBAuth,      // conexao chave valor
		ConnDBEntidades: ConnDBEntidades, // conexao relacional
	}
}
