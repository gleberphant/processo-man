package apresentacao

import (
	"encoding/json"
	"net/http"
)

// todo
func ExibirJsonApi(w http.ResponseWriter, dados interface{}) error {
	// Converte  'dados' para o formato JSON.
	jason, err := json.Marshal(dados)

	if err != nil {
		return err
	}

	// Define o cabeçalho da resposta para indicar que o conteúdo é JSON.
	w.Header().Set("Contet-Type", "application/json")

	// Escreve o JSON resultante no corpo da resposta.
	_, err = w.Write(jason)

	if err != nil {
		return err
	}
	return nil
}
