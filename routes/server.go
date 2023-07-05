package routes

import (
	"os"

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
	app.Get("/Validar", Valida_db)

	port := os.Getenv("HTTP_PLATFORM_PORT")

	// default back to 8080 for local dev
	if port == "" {
		port = "8080"
	}

	app.Listen("127.0.0.1:" + port)

}
