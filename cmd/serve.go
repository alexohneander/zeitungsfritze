package cmd

import (
	"github.com/alexohneander/zeitungsfritze/http"
	"github.com/urfave/cli"
)

func serveCmd(version string) cli.Command {

	return cli.Command{
		Name: "serve",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "port, p",
				Value: "4000",
				Usage: "Port to listen on",
			},
		},
		Action: func(ctx *cli.Context) {
			http.StartServer(version, ctx)
		},
	}
}
