package apresentacao

import (
	"log"
	"net/http"
)

func ExibirErro(w http.ResponseWriter, erroMsg string) {

	log.Println(erroMsg)

	//substituir por redirecionamento para o index com uma mensagem
	http.Error(w, erroMsg, http.StatusInternalServerError)

}
