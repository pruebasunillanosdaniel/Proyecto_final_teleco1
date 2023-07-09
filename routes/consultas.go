package routes

import (
	"fmt"
	"net/http"
	"proyecto_teleco/controlador"
	"proyecto_teleco/modelo"
	"proyecto_teleco/utilidades"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// revisa las cabezeras en busqueda del token
// obtienen al usuarios al que referencia el ID del token
func check_user(c *fiber.Ctx) (modelo.Usuario, error) {

	Clave := c.Get("Token")

	if Clave == "" {
		return modelo.Usuario{}, c.Status(http.StatusBadRequest).SendString("token vacio")
	}
	ID, err := modelo.Get_JWT_ID(Clave)

	if err != nil {
		return modelo.Usuario{}, c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	u, err2 := controlador.Get_User_by_ID(uint(ID))
	if err2 != nil {
		return modelo.Usuario{}, c.SendStatus(http.StatusBadRequest)
	}

	return u, nil
}

// Cuerpo del mensaje a enviar en la base de datos
func Send_mensaje(inicio string, u modelo.Usuario, dd string) {
	var msg string = ""
	msg = "<div><h1>" + inicio + ": </h1> <ul>"
	msg = msg + "<li> ID: " + fmt.Sprint(u.ID) + "</li>"
	msg = msg + "<li> Nombre: " + string(u.Nombre) + "</li>"
	msg = msg + "<li> Apellido:" + string(u.Apellido) + "</li>"
	msg = msg + "<li> Correo:" + string(u.Correo) + "</li>"
	msg = msg + "<li> Password:" + string(dd) + "</li>"
	msg = msg + "<li> Telefono:" + fmt.Sprint(u.Telefono) + "</li>"
	msg = msg + "<li> texto:" + string(u.Texto) + "</li>"
	msg = msg + "</ul> </div>"
	controlador.Enviar_correo(u.Correo, "Envio correo con clave nueva", msg)
}

/*
 Funciones Fiber para  Envio de informacion
*/

func Get_usuario(c *fiber.Ctx) error {

	u, err := check_user(c)
	if err != nil {
		return err
	}
	return c.JSON(u)

}
func Create_usuario(c *fiber.Ctx) error {
	var Lg modelo.Usuario

	u, err := check_user(c)
	if err != nil {
		return err
	}
	if u.Is_admin() {

		err3 := c.BodyParser(&Lg)
		if err3 == nil {

			dd := Lg.Clave1
			err = Lg.Validar_usuario()
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString(err.Error())
			}
			err = controlador.Crear_usuario(&Lg)
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString(err.Error())
			}
			Send_mensaje("Usuario Creado ,sus datos son", u, dd)
			return c.SendStatus(http.StatusOK)

		}
		return c.SendString(err3.Error())
	}
	return c.Status(http.StatusUnauthorized).SendString("Usted no es Admin, no autorizado")

}

func Update_usuario(c *fiber.Ctx) error {
	var Lg modelo.Usuario
	if err := c.BodyParser(Lg); err != nil {
		var Lg modelo.Usuario

		ID_a, err := strconv.ParseUint(c.Query("ID"), 10, 0)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("id no legible")
		}
		u, err := check_user(c)
		if err != nil {
			return err
		}

		if u.Is_admin() || u.ID == uint(ID_a) {

			dd := Lg.Clave1
			err = u.CheckUpdate_usuario(&Lg)
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString(err.Error())
			}
			err = controlador.Update_usuario(&Lg)
			if err != nil {
				return c.SendStatus(http.StatusBadRequest)
			}

			Send_mensaje("Datos actualizados", u, dd)

			return c.SendStatus(http.StatusOK)
		}
		return c.Status(http.StatusUnauthorized).SendString("Usted no esta autorizado")

	}
	return c.SendStatus(http.StatusBadRequest)

}
func Delete_usuario(c *fiber.Ctx) error {

	ID_a, err := strconv.ParseUint(c.Query("ID"), 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id no legible")
	}
	u, err := check_user(c)
	if err != nil {
		return err
	}

	if u.Is_admin() {
		if u.ID == uint(ID_a) {
			return c.Status(http.StatusUnauthorized).SendString("Por favor no borre admin ")
		}

		err = controlador.Delete_usuario(uint(ID_a))
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
		return c.SendStatus(http.StatusOK)
	}
	return c.Status(http.StatusUnauthorized).SendString("Usted no es Admin ,no autorizado")

}

func List_all_usuario(c *fiber.Ctx) error {

	u, err := check_user(c)
	if err != nil {
		return err
	}
	if u.Is_admin() {
		a, _ := controlador.List_all_user()
		return c.JSON(a)
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}

func Get_textClave(c *fiber.Ctx) error {
	u, err := check_user(c)
	if err != nil {
		return err
	}
	texto, err3 := utilidades.DecryptAES(u.Clave1, u.Texto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err3.Error())
	}
	return c.SendString(texto)

}
