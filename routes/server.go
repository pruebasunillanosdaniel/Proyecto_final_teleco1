package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Crear_servidor() {

	var app *fiber.App = fiber.New()

	app.Static("/css", "./vista/src/html/css")
	app.Get("/hola", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"aaa": "asdasdf"})
	})

	app.Get("", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"aaa": "asdasdf"})
	})
	app.Get("/List_users", List_all_usuario)
	app.Post("/crearusuario", Create_usuario)
	app.Get("/leer_usaurio", Get_usuario)
	app.Patch("/actualizarusuario", Update_usuario)
	app.Delete("/eliminar_usuario", Delete_usuario)
	app.Get("/validar", Valida_db)
	app.Get("/paths", Valida_db_variables)

	port := os.Getenv("HTTP_PLATFORM_PORT")

	// default back to 8080 for local dev
	if port == "" {
		port = "80"
	}

	app.Listen(":" + "9000")

}
