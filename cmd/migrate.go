package cmd

import (
	"github.com/alexohneander/zeitungsfritze/database"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

func migrateCmd() cli.Command {

	return cli.Command{
		Name: "migrate",
		Action: func(ctx *cli.Context) {
			database.Connect()
			migrateErr := database.Migrate(database.DB)

			if migrateErr != nil {
				log.Panic().Msg("Error migrating database: " + migrateErr.Error())
			}
		},
	}
}
