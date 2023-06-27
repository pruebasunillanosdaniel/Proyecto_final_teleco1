package controlador

import (
	"crypto/tls"
	"net/smtp"
	"os"
)

var (
	from     string = os.Getenv("usuario_correo")
	password string = os.Getenv("password_origen")
)

const (
	addre string = "smtp.gmail.com"
	port  string = "587"
)

// Funcion auxiliar que envia mensajes a los interesados
func Enviar_mensaje(correo string, subject string, mensaje string) error {

	var user string = from
	var password string = password

	addr := "smtp.gmail.com:465"
	host := "smtp.gmail.com"

	msg := []byte("From:" + from + "\r\n" +
		"To:" + correo + "\r\n" +
		"Subject:" + subject + "\r\n\r\n" +
		"" + mensaje + " \r\n")

	auth := smtp.PlainAuth("", user, password, host)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(from); err != nil {
		return err
	}

	if err = c.Rcpt(correo); err != nil {
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	c.Quit()

	return nil
}
