package appconfig

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetDiretorioRaiz() string {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Printf("error ao encontrar diretorio raiz %v", err)
	}

	for {
		//procura arquivo go.mod
		_, err := os.Stat(filepath.Join(dir, "go.mod"))
		//go.mod encontrado, então diretorio raiz
		if err == nil {
			return dir
		}

		anterior := filepath.Dir(dir)

		if anterior == dir {
			//chegou na raiz do sistema
			log.Printf("\n Diretorio não encontrado. %s \n", dir)
			dir, _ = os.Getwd()
			return dir
		}
		dir = anterior
	}
}
