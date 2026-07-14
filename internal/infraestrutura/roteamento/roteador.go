package roteamento

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/appconfig"
	"github.com/gleberphant/ProcessoMan/internal/dominio/servicos"
)

type IManipulador interface {
	InjetarRotas(mux *http.ServeMux)
	Fechar()
}

// struct que carrega o roteador http  da aplicação
type Roteador struct {
	Handler             http.Handler
	servicoAutenticacao *servicos.ServicoAutenticacao
	manipuladores       []IManipulador
	appConfig           *appconfig.AppConfig
}

func NovoRoteador(appConfig *appconfig.AppConfig) *Roteador {
	return &Roteador{
		appConfig: appConfig,
	}
}

func (r *Roteador) Fechar() {
	for _, m := range r.manipuladores {
		m.Fechar()
	}

}
