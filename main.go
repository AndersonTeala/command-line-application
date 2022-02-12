package main

import (
	"fmt"
	"command-line/app"
	"os"
	"log"
)

func main() {
	fmt.Println("Iniciando......")

	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}

}