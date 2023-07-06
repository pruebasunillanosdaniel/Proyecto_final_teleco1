package routes

import (
	"fmt"
	"proyecto_teleco/database"
	"proyecto_teleco/utilidades"

	"github.com/gofiber/fiber/v2"
)

func Valida_db(c *fiber.Ctx) error {
	db, err := database.Database()

	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})

	}
	return c.JSON(fiber.Map{"error": "OK" + fmt.Sprint(db)})

}

func Valida_db_variables(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"host":     utilidades.Azure_host,
		"port":     utilidades.Azure_port,
		"user":     utilidades.Azure_user,
		"password": utilidades.Azure_password,
		"db":       utilidades.Azure_db,
	})

}
