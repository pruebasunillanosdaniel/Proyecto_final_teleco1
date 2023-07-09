package modelo

import (
	"errors"
	"fmt"
	"proyecto_teleco/utilidades"
)

type Tipo_ide string

const (
	Cedula              Tipo_ide = "CC"
	Extrangeria         Tipo_ide = "CE"
	Documento_identidad Tipo_ide = "TI"
)

type Usuario struct {
	ID       uint     `json:"ID,omitempty" gorm:"primaryKey"`
	Nombre   string   `json:"Nombre"`
	Apellido string   `json:"Apellido"`
	Telefono int      `json:"Telefono"`
	Tipo_id  Tipo_ide `json:"Tipo_id" gorm:"index:idx_member,unique"`
	Num_ide  uint     `json:"Num_ide" gorm:"index:idx_member,unique"`
	Admin    bool     `json:"admin,omitempty" `
	Texto    string   `json:"Texto"`
	Correo   string   `json:"Correo"`
	Clave1   string   `json:"Clave"`
	TokenID  uint     `json:"-"`
}

func (U *Usuario) Is_admin() bool {
	return U.Admin
}
func (U *Usuario) Set_admin() {
	U.Admin = true
}

func (U *Usuario) Validar_llave(clave string) error {
	var nueva_clave string = utilidades.GenerarSHA256_with_32bits(clave)
	if nueva_clave == U.Clave1 {
		return nil
	}
	return errors.New("error,llave incorrecta ")
}

func Validaciones(U *Usuario) error {
	if utilidades.Validar_correo(U.Correo) {
		return errors.New("error Correo Incorrecto")
	}
	if utilidades.Validar_telefono(fmt.Sprint(U.Telefono)) {
		return errors.New("error: Telefono Incorrecto ,porfavor un numero de exactamente 10 digitos")
	}
	if utilidades.Validar_clave_signos(U.Correo) {
		return errors.New("error Correo Incorrecto")
	}
	if U.Texto == "" {
		return errors.New("error Texto  Vacio")
	}
	if U.Clave1 != "" {
		return errors.New("error Clave  Vacio")
	}
	return nil

}

func (U *Usuario) Validar_usuario() error {

	if errv := Validaciones(U); errv != nil {
		return errv
	}
	clave2 := utilidades.GenerarSHA256_with_32bits(U.Clave1)

	enc, err := utilidades.EncryptAES(clave2, U.Texto)
	if err != nil {
		return errors.New("error, Clave no pudo ser creado")
	}
	U.Texto = enc

	return nil
}

func (U *Usuario) CheckUpdate_usuario(Un *Usuario) error {
	if errv := Validaciones(Un); errv != nil {
		return errv
	}

	if U.Clave1 != Un.Clave1 {

		txt, err2 := utilidades.DecryptAES(U.Clave1, U.Texto)
		if err2 != nil {
			return err2
		}
		clave2 := utilidades.GenerarSHA256_with_32bits(Un.Clave1)
		enc, err := utilidades.EncryptAES(clave2, txt)
		if err != nil {
			return err
		}
		Un.Texto = enc
		return nil
	}

	return errors.New("Error , porfavor enviar los datos correctamente ,")

}
