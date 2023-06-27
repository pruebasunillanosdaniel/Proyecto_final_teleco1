package modelo

import (
	"errors"
	"proyecto_teleco/utilidades"
)

type Tipo_ide string

const (
	Cedula              Tipo_ide = "CC"
	Extrangeria         Tipo_ide = "CE"
	Documento_identidad Tipo_ide = "TI"
)

type Usuario struct {
	Id       int
	Nombre   string
	Apellido string
	Telefono int
	Tipo_id  Tipo_ide
	Cedula   int
	Texto    string
	Correo   string
	clave1   string
	clave2   string
}

func (U *Usuario) Validar_llave(clave string) error {

	var nueva_clave string = utilidades.GenerarSHA254(clave)
	if nueva_clave == U.clave1 {
		return nil
	}
	return errors.New("Error,llave incorrecta ")
}

func Crear_usuario(Nombre string, Apellido string, Telefono int, Cedula int, Texto string, Correo, Clave string) error {

	if utilidades.Validar_correo(Correo) {
		return errors.New("Error Correo Incorrecto")
	}

	var u Usuario = Usuario{
		Nombre:   Nombre,
		Apellido: Apellido,
		Cedula:   Cedula,
	}

}
