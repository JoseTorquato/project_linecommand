package main

import (
	"fmt"
	"linha_comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Iniciando o projeto")

	// instanciando a aplicação no main
	aplicacao := app.Gerar()
	// if init para erro simples
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
