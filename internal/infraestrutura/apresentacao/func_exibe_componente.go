package apresentacao

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

// com essa função eu posso exibir um componente sem precisar carregar TODOs os componentes no template de cada pagina
func exibeComponente(chave string, dados interface{}) template.HTML {
	var buffer bytes.Buffer

	componente, ok := cacheComponentes[chave]

	if !ok {
		return htmlErrorBox("Componente <strong>" + chave + "</strong> não encontrado no cache")
	}
	err := componente.Execute(&buffer, dados)

	if err != nil {
		log.Printf("%v", err)
		return htmlErrorBox("Erro ao renderizar o componente <strong>" + chave + "</strong></div>")

	}

	return template.HTML(buffer.String())
}

func htmlErrorBox(mensagem string) template.HTML {
	return template.HTML(
		fmt.Sprintf(
			`<div class="alert alert-danger" role="alert"> %s </div>`, mensagem,
		),
	)
}
