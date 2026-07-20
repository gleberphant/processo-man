package apresentacao

import (
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

func ExibirErro(w http.ResponseWriter, erroMsg string) {

	pc, file, _, _ := runtime.Caller(1)

	//logo no console
	functionName := runtime.FuncForPC(pc).Name()
	file = filepath.Base(file)

	log.Printf("\n *FUNC %s \n *FILE %s \n *ERROR: %s \n", functionName, file, erroMsg)

	//substituir por redirecionamento para o index com uma mensagem
	if w != nil {
		http.Error(w, erroMsg, http.StatusInternalServerError)
	}
}
