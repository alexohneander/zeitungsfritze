package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

func Run(version, commitHash string) {
	app := cli.App{
		Name:    "zeitungsfritze",
		Version: fmt.Sprintf("%s; zeitungsfritze/server@%s", version, commitHash),
		Commands: []cli.Command{
			serveCmd(version),
			migrateCmd(),
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("app error")
	}
}
