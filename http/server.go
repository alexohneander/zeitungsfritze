package http

import (
	"github.com/alexohneander/zeitungsfritze/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

func StartServer(version string, ctx *cli.Context) {
	app := fiber.New(fiber.Config{
		ServerHeader:          "Zeitungsfritze/" + version,
		AppName:               "Zeitungsfritze",
		DisableStartupMessage: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(helper.CreateDomain())
	})

	log.Info().Msg("Starting http server on :" + ctx.String("port"))
	app.Listen(":" + ctx.String("port"))
}
