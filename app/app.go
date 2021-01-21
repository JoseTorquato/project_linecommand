package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servodores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",       // Nome da flag
			Value: "google.com", // padrão caso não seja especificado algum valor
		},
	}

	// Slice de commandos
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca de IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca os servidores que estão hospedados",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

// Função que busca ips e lista os mesmos
func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	// Tratando o erro com log
	if erro != nil {
		log.Fatal(erro)
	}

	// Por ser um slice eu posso iterar em ips
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	// Tratando o erro com log
	if erro != nil {
		log.Fatal(erro)
	}
	// Por ser um slice eu posso iterar em servidores
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}
