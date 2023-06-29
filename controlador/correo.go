package controlador

import (
	"proyecto_teleco/utilidades"

	"gopkg.in/gomail.v2"
)

func Enviar_correo(correo string, Subject string, mensaje string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", utilidades.Correo_autor)
	m.SetHeader("To", correo)
	m.SetHeader("Subject", Subject)
	m.SetBody("text/html", mensaje)
	d := gomail.NewDialer(utilidades.Smtp_autor, 587, utilidades.Correo_autor, utilidades.Password_autor)
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
