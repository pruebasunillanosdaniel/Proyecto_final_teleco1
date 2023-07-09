package routes

import (
	"fmt"
	"proyecto_teleco/controlador"
	"proyecto_teleco/database"
	"proyecto_teleco/modelo"
	"proyecto_teleco/utilidades"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	var Clave string = c.Get("Clave")

	if Clave != "" {
		return c.Status(fiber.StatusBadRequest).SendString("Clave no reconocible")
	}

	UserID, err := strconv.ParseUint(c.Get("UserID"), 10, 0)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Id no reconocible")
	}
	u, err := controlador.Get_User_by_ID(uint(UserID))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := u.Validar_llave(Clave); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	JT, err := controlador.Read_JWT_usuario(u.ID)
	if err == nil && !JT.Is_valid() {
		controlador.Delete_JWT(JT.ID)
	} else {
		return c.Status(fiber.StatusOK).SendString(JT.Datos)
	}
	JWT_token, err := modelo.Create_Jwt_database(u)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	controlador.Crear_JWT(&JWT_token)
	return c.Status(fiber.StatusOK).SendString(JWT_token.Datos)

}

func Logout(c *fiber.Ctx) error {

	u, err := check_user(c)
	if err != nil {
		return err
	}
	err = controlador.Delete_JWT_usuario(u.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).SendString("Gracias por su visita")

}

func Valida_db(c *fiber.Ctx) error {
	db, err := database.Database()

	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})

	}
	return c.JSON(fiber.Map{"OK": "OK" + fmt.Sprint(db)})

}

func Valida_db_variables(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"host": utilidades.Azure_host,
		"port": utilidades.Azure_port,
		"db":   utilidades.Azure_db,
	})

}
