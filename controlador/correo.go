package controlador

import (
	"net/smtp"
	"os"
	"cripto/tls"
)

var(
	user_origen     string = os.Getenv("usuario_correo")
	password_origen string = os.Getenv("password_origen")
)

const (
	addre string = "smtp.gmail.com"
	port  string = "587"
	tlsconfig := &tls.tlsconfig{

		 InsecureSkipVerify: true,
        	 ServerName: host,
	}
)



// Funcion auxiliar que envia mensajes a los interesados
func Enviar_mensaje(correo string, mensaje string) {
	
	var bmensaje []byte = []byte(mensaje)
	
	auth:=smtp.PlainAuth("")

}
