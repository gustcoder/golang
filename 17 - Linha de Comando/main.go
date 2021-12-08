package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	// usando if init
	if err := aplicacao.Run(os.Args); err != nil {
		// Fatal para a aplicacao
		log.Fatal(err)
	}
}
