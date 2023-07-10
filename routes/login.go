package routes

import (
	"fmt"
	"log"
	"proyecto_teleco/controlador"
	"proyecto_teleco/database"
	"proyecto_teleco/modelo"
	"proyecto_teleco/utilidades"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	var Clave string = c.Get("Clave")

	if Clave == "" {
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
	log.Println("-- log mensaje read jwt", err)

	if err != nil {
		log.Println("crear nuevo token")
		JWT_token, err := modelo.Create_Jwt_database(u)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		JWT_token.Token = u.ID
		err = controlador.Crear_JWT(&JWT_token)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.Status(fiber.StatusOK).SendString(JWT_token.Datos)
	} else {
		log.Println(" token antiguo", JT.Datos, JT.Is_valid())

		if !JT.Is_valid() {
			err = controlador.Delete_JWT(JT.ID)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			JWT_token, err := modelo.Create_Jwt_database(u)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
			JWT_token.Token = u.ID
			err = controlador.Crear_JWT(&JWT_token)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}

			return c.Status(fiber.StatusOK).SendString(JWT_token.Datos)

		}

		return c.Status(fiber.StatusOK).SendString(JT.Datos)

	}
}

func Logout(c *fiber.Ctx) error {

	u, err := check_user(c)
	log.Println("objetos de logout", u, err)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	log.Println(u, err)
	err = controlador.Delete_JWT_usuario(u.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
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
