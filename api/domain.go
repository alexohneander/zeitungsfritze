package api

import (
	"github.com/alexohneander/zeitungsfritze/database"
	"github.com/alexohneander/zeitungsfritze/helper"
	"github.com/alexohneander/zeitungsfritze/model"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

var db, dbErr = database.Connect()

func CreateDomain(ctx *fiber.Ctx) error {
	domain := helper.CreateDomain()
	forward := model.Forward{Domain: domain, Target: "127.0.0.1", Port: "8080"}

	result := db.Create(&forward)

	if result.Error != nil || dbErr != nil {
		log.Error().Msg("Error creating domain: " + result.Error.Error())

		ctx.Status(500)
		return ctx.JSON(fiber.Map{"status": "error"})
	}

	return ctx.JSON(fiber.Map{"status": "success", "domain": domain})
}

func DeleteDomain(ctx *fiber.Ctx) error {
	result := db.Where("domain = ?", ctx.Params("domain")).Delete(&model.Forward{})

	if result.Error != nil || dbErr != nil {
		log.Error().Msg("Error deleting domain: " + result.Error.Error())

		ctx.Status(500)
		return ctx.JSON(fiber.Map{"status": "error"})
	}

	return ctx.JSON(fiber.Map{"status": "success"})
}

func GetDomain() {

}

func GetDomains() {

}

func UpdateDomain() {

}
