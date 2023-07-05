package routes

import (
	"proyecto_teleco/database"

	"github.com/gofiber/fiber/v2"
)

func Valida_db(c *fiber.Ctx) error {
	_, err := database.Database()
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})

	}
	return c.JSON(fiber.Map{"error": "OK"})

}
