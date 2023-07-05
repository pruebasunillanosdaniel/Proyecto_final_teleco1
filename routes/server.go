package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Crear_servidor() {

	var app *fiber.App = fiber.New()

	app.Static("/css", "./vista/src/html/css")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"aaa": "asdasdf"})
	})

	app.Get("/List_users", List_all_usuario)
	app.Post("/Create", Create_usuario)
	app.Get("/READ/", Get_usuario)
	app.Patch("/UPDATE", Update_usuario)
	app.Delete("/List_users", Delete_usuario)
	app.Listen(":8080")

}
