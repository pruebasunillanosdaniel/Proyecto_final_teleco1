package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Crear_servidor() {

	var app *fiber.App = fiber.New()

	app.Static("/css", "./vista/src/html/css")
	app.Get("/hola", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"Saludo": "SAludos a todos esto es por >DEBUG"})
	})

	app.Get("", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"aaa": "asdasdf"})
	})
	app.Get("/login", Login)
	app.Get("/List_users", List_all_usuario)
	app.Post("/Create_user", Create_usuario)
	app.Get("/Read_user", Get_usuario)
	app.Put("/Update_user", Update_usuario)
	app.Delete("/Delete", Delete_usuario)
	app.Get("/Text", Get_textClave)
	app.Get("/DB_connection", Valida_db)
	app.Get("/logout", Logout)

	//app.Get("/paths", Valida_db_variables)

	port := os.Getenv("HTTP_PLATFORM_PORT")

	// default back to 8080 for local dev
	if port == "" {
		port = "8080"
	}

	app.Listen(":" + port)

}
