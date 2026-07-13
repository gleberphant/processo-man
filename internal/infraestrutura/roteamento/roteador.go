package roteamento

import (
	"database/sql"
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/dominio/servicos"
	"go.etcd.io/bbolt"
)

type IManipulador interface {
	InjetarRotas(mux *http.ServeMux)
	Fechar()
}

type DBConfig struct {
	ConnDBAuth      *bbolt.DB // conexao chave valor
	ConnDBEntidades *sql.DB   // conexao relacional
}

type Roteador struct {
	Handler             http.Handler
	servicoAutenticacao *servicos.ServicoAutenticacao
	manipuladores       []IManipulador
	dbConfig            *DBConfig
}

func NovoRoteador(db *DBConfig) *Roteador {
	return &Roteador{

		dbConfig: db,
	}
}

func (r *Roteador) Fechar() {
	for _, m := range r.manipuladores {
		m.Fechar()
	}

}
