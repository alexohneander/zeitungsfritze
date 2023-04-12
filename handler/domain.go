package handler

import (
	"github.com/alexohneander/zeitungsfritze/database"
	"github.com/alexohneander/zeitungsfritze/helper"
	"github.com/alexohneander/zeitungsfritze/model"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func CreateDomain(ctx *fiber.Ctx) error {
	db := database.DB
	domain := helper.CreateDomain()
	forward := model.Forward{Domain: domain, Target: "127.0.0.1", Port: "8080"}

	result := db.Create(&forward)

	if result.Error != nil {
		log.Error().Msg("Error creating domain: " + result.Error.Error())

		ctx.Status(500)
		return ctx.JSON(fiber.Map{"status": "error"})
	}

	log.Info().Msg("Created domain: " + domain)
	return ctx.JSON(fiber.Map{"status": "success", "domain": domain})
}

func DeleteDomain(ctx *fiber.Ctx) error {
	db := database.DB
	result := db.Where("domain = ?", ctx.Params("domain")).Delete(&model.Forward{})

	if result.Error != nil {
		log.Error().Msg("Error deleting domain: " + result.Error.Error())

		ctx.Status(500)
		return ctx.JSON(fiber.Map{"status": "error"})
	}

	log.Info().Msg("Deleted domain: " + ctx.Params("domain"))
	return ctx.JSON(fiber.Map{"status": "success"})
}

func GetDomain(ctx *fiber.Ctx) error {
	db := database.DB
	var forward model.Forward

	result := db.Where("domain = ?", ctx.Params("domain")).First(&forward)

	if result.Error != nil {
		log.Error().Msg("Error getting domain: " + result.Error.Error())

		ctx.Status(500)
		return ctx.JSON(fiber.Map{"status": "error"})
	}

	log.Info().Msg("Got domain: " + ctx.Params("domain"))
	return ctx.JSON(fiber.Map{"status": "success", "domain": forward})
}

func GetDomains(ctx *fiber.Ctx) error {
	db := database.DB
	var forwards []model.Forward

	result := db.Find(&forwards)

	if result.Error != nil {
		log.Error().Msg("Error getting domains: " + result.Error.Error())

		ctx.Status(500)
		return ctx.JSON(fiber.Map{"status": "error"})
	}

	log.Info().Msg("Got domains")
	return ctx.JSON(fiber.Map{"status": "success", "domains": forwards})
}

func UpdateDomain(ctx *fiber.Ctx) error {
	db := database.DB
	var forward model.Forward

	result := db.Where("domain = ?", ctx.Params("domain")).First(&forward)

	if result.Error != nil {
		log.Error().Msg("Error updating domain: " + result.Error.Error())

		ctx.Status(500)
		return ctx.JSON(fiber.Map{"status": "error"})
	}

	result = db.Model(&forward).Updates(model.Forward{Target: ctx.FormValue("target"), Port: ctx.FormValue("port")})

	if result.Error != nil {
		log.Error().Msg("Error updating domain: " + result.Error.Error())

		ctx.Status(500)
		return ctx.JSON(fiber.Map{"status": "error"})
	}

	log.Info().Msg("Updated domain: " + ctx.Params("domain"))
	return ctx.JSON(fiber.Map{"status": "success"})
}
