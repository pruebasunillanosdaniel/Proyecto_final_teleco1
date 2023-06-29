package modelo

import (
	"errors"
	"proyecto_teleco/utilidades"

	"gorm.io/gorm"
)

type Tipo_ide string

const (
	Cedula              Tipo_ide = "CC"
	Extrangeria         Tipo_ide = "CE"
	Documento_identidad Tipo_ide = "TI"
)

type Usuario struct {
	gorm.Model
	Nombre   string
	Apellido string
	Telefono int
	Tipo_id  Tipo_ide `gorm:"index:idx_member"`
	Num_ide  int      `gorm:"index:idx_member"`
	Texto    string
	Correo   string
	clave1   string
}

func (U *Usuario) Validar_llave(clave string) error {

	var nueva_clave string = utilidades.GenerarSHA254(clave)
	if nueva_clave == U.clave1 {
		return nil
	}
	return errors.New("error,llave incorrecta ")
}

func Crear_usuario(Nombre string, Apellido string, Telefono int,
	Num_ide int, Texto string, Correo, Clave string) (error, Usuario) {

	if utilidades.Validar_correo(Correo) {
		return errors.New("error Correo Incorrecto"), Usuario{}
	}
	clave2 := utilidades.GenerarSHA254(Clave)

	enc, err := utilidades.EncryptAES([]byte(clave2), Texto)
	if err != nil {

		return errors.New("error, Clave no pudo ser creado"), Usuario{}
	}

	var u Usuario = Usuario{
		Nombre:   Nombre,
		Apellido: Apellido,
		Num_ide:  Num_ide,
		Texto:    enc,
		clave1:   clave2,
	}
	return nil, u
}
