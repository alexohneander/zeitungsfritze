package http

import (
	"github.com/alexohneander/zeitungsfritze/api"
	"github.com/gofiber/fiber/v2"
)

func ConfigureRoutes(app *fiber.App) *fiber.App {
	apiGroup := app.Group("/api")
	v1Group := apiGroup.Group("/v1")

	v1Group.Get("/domains/create", api.CreateDomain) // /api/v1/domains/create

	return app
}
