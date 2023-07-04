package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Crear_servidor() {

	var app *fiber.App = fiber.New()

	app.Static("/css", "./vista/src/html/css")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("vista/src/html/index.html", fiber.Map{})
	})
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("vista/src/html/registrar.html", fiber.Map{})
	})

	app.Listen(":8080")
}
