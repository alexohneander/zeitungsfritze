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
			db, connectErr := database.Connect()
			migrateErr := database.Migrate(db)

			if connectErr != nil || migrateErr != nil {
				log.Error().Msg("Error migrating database: " + connectErr.Error() + migrateErr.Error())
			}
		},
	}
}
