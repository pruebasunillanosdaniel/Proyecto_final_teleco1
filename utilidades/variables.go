package utilidades

import "os"

var (
	Correo_autor   string = os.Getenv("correo_usuario")
	Smtp_autor     string = os.Getenv("smtp_server")
	Password_autor string = os.Getenv("password_usuario")
	Azure_host     string = os.Getenv("host_db_azure")
	Azure_db       string = os.Getenv("db_azure")
	Azure_port     string = os.Getenv("port_azure")
	Azure_user     string = os.Getenv("user_db_user")
	Azure_password string = os.Getenv("azure_password")
	Secret_jwt     string = os.Getenv("secret_jwt")
	Clave_texto16  string = os.Getenv("clave_texto16")
)

//teleinfouser.database.windows.net
