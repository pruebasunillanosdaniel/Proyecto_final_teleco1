package routes

import (
	"net/http"
	"proyecto_teleco/controlador"
	"proyecto_teleco/modelo"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get usuario de la base de datos
// @param ID
func Get_usuario(c *fiber.Ctx) error {
	var Lg modelo.Loguin_persona
	if err := c.BodyParser(&Lg); err != nil {
		u, err2 := controlador.Get_User_by_unique(Lg.Tipo_ide, Lg.Numero_identidad)
		if err2 != nil {
			return c.SendStatus(http.StatusBadRequest)
		}
		if u.Validar_llave(Lg.Password) != nil {
			return c.SendStatus(http.StatusUnauthorized)
		}
		return c.JSON(u)
	}
	return c.SendStatus(http.StatusBadRequest)
}
func Create_usuario(c *fiber.Ctx) error {
	var Lg modelo.Login_Datos
	if err := c.BodyParser(Lg); err != nil {
		u, err2 := controlador.Get_User_by_unique(Lg.Login.Tipo_ide, Lg.Login.Numero_identidad)
		if err2 != nil {
			return c.Status(http.StatusBadRequest).SendString(err2.Error())
		}

		if u.Is_admin() {
			if u.Validar_llave(Lg.Login.Password) != nil {
				return c.Status(http.StatusUnauthorized).SendString("llave incorrecta ")
			}
			err = Lg.Datos.Validar_usuario()
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString(err.Error())
			}
			err = controlador.Crear_usuario(&Lg.Datos)
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString(err.Error())
			}
			return c.SendStatus(http.StatusOK)
		}
		return c.Status(http.StatusUnauthorized).SendString("Usted no es Admin, no autorizado")

	}
	return c.SendStatus(http.StatusBadRequest)

}

func Update_usuario(c *fiber.Ctx) error {
	var Lg modelo.Login_Datos
	if err := c.BodyParser(Lg); err != nil {
		u, err2 := controlador.Get_User_by_unique(Lg.Login.Tipo_ide, Lg.Login.Numero_identidad)
		if err2 != nil {
			return c.SendStatus(http.StatusBadRequest)
		}
		if u.Validar_llave(Lg.Login.Password) != nil {
			return c.Status(http.StatusUnauthorized).SendString("llave incorrecta ")
		}
		if u.Is_admin() || (Lg.Datos.Num_ide == u.Num_ide && Lg.Datos.Tipo_id == u.Tipo_id) {
			err = Lg.Datos.Validar_usuario()
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString(err.Error())
			}
			err = controlador.Update_usuario(&Lg.Datos)
			if err != nil {
				return c.SendStatus(http.StatusBadRequest)
			}
			return c.SendStatus(http.StatusOK)
		}
		return c.Status(http.StatusUnauthorized).SendString("Usted no esta autorizado")

	}
	return c.SendStatus(http.StatusBadRequest)

}
func Delete_usuario(c *fiber.Ctx) error {
	var Lg modelo.Loguin_persona
	if err := c.BodyParser(Lg); err != nil {
		u, err2 := controlador.Get_User_by_unique(Lg.Tipo_ide, Lg.Numero_identidad)
		if err2 != nil {
			return c.Status(fiber.StatusBadRequest).SendString("error ::: " + err.Error())
		}

		if u.Is_admin() {
			if u.Validar_llave(Lg.Password) != nil {
				return c.Status(http.StatusUnauthorized).SendString("llave incorrecta ")
			}
			ID, err3 := strconv.ParseUint(Lg.ID, 10, -1)
			if err3 != nil {
				return c.SendStatus(http.StatusBadRequest)
			}
			err = controlador.Delete_usuario(uint(ID))
			if err != nil {
				return c.SendStatus(http.StatusBadRequest)
			}
			return c.SendStatus(http.StatusOK)
		}
		return c.Status(http.StatusUnauthorized).SendString("Usted no es Admin no autorizado")

	}
	return c.SendStatus(http.StatusBadRequest)

}

func List_all_usuario(c *fiber.Ctx) error {
	var Lg modelo.Loguin_persona
	if err := c.BodyParser(Lg); err != nil {
		u, err2 := controlador.Get_User_by_unique(Lg.Tipo_ide, Lg.Numero_identidad)
		if err2 != nil {
			return c.Status(fiber.StatusBadRequest).SendString("error ::: " + err.Error())
		}
		if u.Is_admin() {
			a, _ := controlador.List_all_user()
			return c.JSON(a)
		}
	}
	return c.SendStatus(http.StatusBadRequest)

}
