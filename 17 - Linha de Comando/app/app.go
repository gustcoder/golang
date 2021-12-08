package app

import "github.com/urfave/cli"

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App Linha de Comando"
	app.Usage = "Busca IPs e nomes de servidores"

	return app
}
