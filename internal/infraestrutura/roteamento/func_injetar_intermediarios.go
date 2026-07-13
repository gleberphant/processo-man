package roteamento

import (
	"net/http"

	"github.com/gleberphant/ProcessoMan/internal/aplicacao/intermediarios"
)

func (r *Roteador) InjetarIntermediarios() {

	var roteador http.Handler
	// INJETA INTERMEDIÁRIOS - Middlewares
	roteador = intermediarios.AutenticadorFunc(r.Handler, r.servicoAutenticacao)
	roteador = intermediarios.LoggerFunc(roteador)

	r.Handler = roteador

}
