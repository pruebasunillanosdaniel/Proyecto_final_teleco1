package utilidades

import "os"

var (
	Correo_autor   string = os.Getenv("correo_usuario")
	Smtp_autor     string = os.Getenv("smtp_server")
	Password_autor string = os.Getenv("password_usuario")
	Azure_host     string = os.Getenv("azure_host")
	Azure_db       string = os.Getenv("azure_db")
	Azure_port     string = os.Getenv("azure_port")
	Azure_user     string = os.Getenv("azure_user")
	Azure_password string = os.Getenv("azure_password")
	Secret_jwt     string = os.Getenv("secret_jwt")
	Clave_texto16  string = os.Getenv("clave_texto16")
)

//teleinfouser.database.windows.net
