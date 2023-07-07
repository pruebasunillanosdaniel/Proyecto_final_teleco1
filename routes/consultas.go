package routes

import (
	"log"
	"net/http"
	"proyecto_teleco/controlador"
	"proyecto_teleco/modelo"
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
	if u.Validar_llave(u.Clave1) != nil {
		return c.SendStatus(http.StatusUnauthorized)
	}
	return c.JSON(u)

}
func Create_usuario(c *fiber.Ctx) error {
	var Lg modelo.Login_Datos
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
		err3 := c.BodyParser(&Lg)
		if err3 == nil {
			if u.Validar_llave(Clave) != nil {
				return c.Status(http.StatusUnauthorized).SendString("llave incorrecta ")
			}
			dd := Lg.Datos.Clave1
			err = Lg.Datos.Validar_usuario()
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString(err.Error())
			}
			err = controlador.Crear_usuario(&Lg.Datos)
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString(err.Error())
			}
			var msg string = ""
			msg = "<div><h1>Datos </h1> <ul>"
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
		return c.SendString(err3.Error())
	}
	return c.Status(http.StatusUnauthorized).SendString("Usted no es Admin, no autorizado")

}

func Update_usuario(c *fiber.Ctx) error {
	var Lg modelo.Login_Datos
	var msg string = ""
	if err := c.BodyParser(Lg); err != nil {
		var Lg modelo.Login_Datos
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
		if u.Validar_llave(Lg.Login.Password) != nil {
			return c.Status(http.StatusUnauthorized).SendString("llave incorrecta ")
		}
		if u.Is_admin() || u.ID == uint(ID_a) {
			dd := Lg.Datos.Clave1
			err = Lg.Datos.Validar_usuario()
			if err != nil {
				return c.Status(http.StatusBadRequest).SendString(err.Error())
			}
			err = controlador.Update_usuario(&Lg.Datos)
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
	log.Println("admin:::", u.Nombre, u.Is_admin())
	if u.Is_admin() {
		a, _ := controlador.List_all_user()
		return c.JSON(a)
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}
