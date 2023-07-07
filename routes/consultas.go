package routes

import (
	"fmt"
	"log"
	"net/http"
	"proyecto_teleco/controlador"
	"proyecto_teleco/modelo"
	"proyecto_teleco/utilidades"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Get usuario de la base de datos
// @param ID
func Get_usuario(c *fiber.Ctx) error {

	ID, err := strconv.ParseUint(c.Get("UserId"), 10, -1)
	Clave := c.Get("Clave")
	if err != nil || Clave == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	u, err2 := controlador.Get_User_by_ID(uint(ID))
	if err2 != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	if u.Validar_llave(Clave) != nil {
		return c.SendStatus(http.StatusUnauthorized)
	}
	return c.JSON(u)

}
func Create_usuario(c *fiber.Ctx) error {
	var Lg modelo.Usuario
	ID, err := strconv.ParseUint(c.Get("UserId"), 10, 0)
	Clave := c.Get("Clave")

	if err != nil || Clave == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	u, err2 := controlador.Get_User_by_ID(uint(ID))

	if err2 != nil {
		return c.Status(http.StatusBadRequest).SendString(err2.Error())
	}
	if u.Is_admin() {

		if u.Validar_llave(Clave) != nil {
			return c.Status(http.StatusUnauthorized).SendString("llave incorrecta ")
		}
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
			var msg string = ""
			msg = "<div><h1>Usuario Creado ,sus datos son: </h1> <ul>"
			msg = msg + "<li> ID: " + fmt.Sprint(Lg.ID) + "</li>"
			msg = msg + "<li> Nombre: " + string(Lg.Nombre) + "</li>"
			msg = msg + "<li> Apellido:" + string(Lg.Apellido) + "</li>"
			msg = msg + "<li> Correo:" + string(Lg.Correo) + "</li>"
			msg = msg + "<li> Password:" + string(dd) + "</li>"
			msg = msg + "<li> Telefono:" + fmt.Sprint(Lg.Telefono) + "</li>"
			msg = msg + "<li> texto:" + string(Lg.Texto) + "</li>"
			msg = msg + "</ul> </div>"
			controlador.Enviar_correo(u.Correo, "Envio correo con clave nueva", msg)

			return c.SendStatus(http.StatusOK)

		}
		return c.SendString(err3.Error())
	}
	return c.Status(http.StatusUnauthorized).SendString("Usted no es Admin, no autorizado")

}

func Update_usuario(c *fiber.Ctx) error {
	var Lg modelo.Login_Datos
	var msg string = ""
	if err := c.BodyParser(Lg); err != nil {
		var Lg modelo.Usuario
		ID, err := strconv.ParseUint(c.Get("UserId"), 10, 0)
		Clave := c.Get("Clave")
		ID_a, err := strconv.ParseUint(c.Query("ID"), 10, 0)

		if err != nil || Clave == "" {
			return c.SendStatus(http.StatusBadRequest)
		}
		u, err2 := controlador.Get_User_by_ID(uint(ID))
		if err2 != nil {
			return c.Status(http.StatusBadRequest).SendString(err2.Error())
		}

		if u.Validar_llave(Clave) != nil {
			return c.Status(http.StatusUnauthorized).SendString("llave incorrecta ")
		}
		if u.Is_admin() || u.ID == uint(ID_a) {
			dd := Lg.Clave1
			err = Lg.Validar_usuario()
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString(err.Error())
			}
			err = controlador.Update_usuario(&Lg)
			if err != nil {
				return c.SendStatus(http.StatusBadRequest)
			}
			msg = "<div><h1>Datos actualizados </h1> <ul>"
			msg = msg + "<li> ID: " + string(u.ID) + "</li>"
			msg = msg + "<li> Nombre: " + string(u.Nombre) + "</li>"
			msg = msg + "<li> Apellido:" + string(u.Apellido) + "</li>"
			msg = msg + "<li> Correo:" + string(u.Correo) + "</li>"
			msg = msg + "<li> Password:" + string(dd) + "</li>"
			msg = msg + "<li> Telefono:" + string(u.Telefono) + "</li>"
			msg = msg + "<li> texto:" + string(u.Texto) + "</li>"
			msg = msg + "</ul> </div>"
			controlador.Enviar_correo(u.Correo, "Envio correo con clave nueva", msg)
			return c.SendStatus(http.StatusOK)
		}
		return c.Status(http.StatusUnauthorized).SendString("Usted no esta autorizado")

	}
	return c.SendStatus(http.StatusBadRequest)

}
func Delete_usuario(c *fiber.Ctx) error {

	ID, err := strconv.ParseUint(c.Get("UserId"), 10, 0)
	Clave := c.Get("Clave")
	ID_a, err := strconv.ParseUint(c.Query("ID"), 10, 0)

	if err != nil || Clave == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	u, err2 := controlador.Get_User_by_ID(uint(ID))
	if err2 != nil {
		return c.Status(http.StatusBadRequest).SendString(err2.Error())
	}

	if u.Is_admin() {
		if u.ID == uint(ID_a) {
			return c.Status(http.StatusUnauthorized).SendString("Por favor no borre admin ")
		}

		if u.Validar_llave(Clave) != nil {
			return c.Status(http.StatusUnauthorized).SendString("llave incorrecta ")
		}

		err = controlador.Delete_usuario(uint(ID_a))
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}
		return c.SendStatus(http.StatusOK)
	}
	return c.Status(http.StatusUnauthorized).SendString("Usted no es Admin no autorizado")

}

func List_all_usuario(c *fiber.Ctx) error {
	log.Println("::::::::::::::...", c.Get("UserId"), c.Get("Clave"))
	ID, err := strconv.ParseUint(c.Get("UserId"), 10, 0)
	Clave := c.Get("Clave")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("error ::: " + err.Error())
	}
	if err != nil || Clave == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	u, err2 := controlador.Get_User_by_ID(uint(ID))
	if err2 != nil {
		return c.Status(fiber.StatusBadRequest).SendString("error ::: " + err.Error())
	}
	if u.Is_admin() {

		log.Println("lista todos usuario clave si ", u.Validar_llave(Clave), Clave, "::::", u.Clave1, Clave == u.Clave1)

		if err4 := u.Validar_llave(Clave); err4 != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err4.Error())
		}
		a, _ := controlador.List_all_user()
		return c.JSON(a)
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}

func Get_textClave(c *fiber.Ctx) error {
	ID, err := strconv.ParseUint(c.Get("UserId"), 10, 0)
	Clave := c.Get("Clave")

	if err != nil || Clave == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	u, err2 := controlador.Get_User_by_ID(uint(ID))

	if err2 != nil {
		return c.Status(http.StatusBadRequest).SendString(err2.Error())
	}

	if err4 := u.Validar_llave(Clave); err4 != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err4.Error())
	}
	texto, err3 := utilidades.DecryptAES(u.Clave1, u.Texto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err3.Error())
	}
	return c.SendString(texto)

}
