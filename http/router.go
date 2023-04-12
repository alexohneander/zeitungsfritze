package http

import (
	"github.com/alexohneander/zeitungsfritze/handler"
	"github.com/alexohneander/zeitungsfritze/middleware"
	"github.com/gofiber/fiber/v2"
)

func ConfigureRoutes(app *fiber.App) *fiber.App {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/login", handler.Login) // /api/v1/auth/login

	// Domains
	domains := v1.Group("/domains")
	domains.Get("/", middleware.Protected(), handler.GetDomains)             // /api/v1/domains
	domains.Post("/", middleware.Protected(), handler.CreateDomain)          // /api/v1/domains
	domains.Get("/:domain", middleware.Protected(), handler.GetDomain)       // /api/v1/domains/:domain
	domains.Patch("/:domain", middleware.Protected(), handler.UpdateDomain)  // /api/v1/domains/:domain
	domains.Delete("/:domain", middleware.Protected(), handler.DeleteDomain) // /api/v1/domains/:domain

	// User
	user := v1.Group("/user")
	user.Post("/", handler.CreateUser)                              // /api/v1/user
	user.Get("/:id", handler.GetUser)                               // /api/v1/user/:id
	user.Patch("/:id", middleware.Protected(), handler.UpdateUser)  // /api/v1/user/:id
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser) // /api/v1/user/:id

	return app
}
