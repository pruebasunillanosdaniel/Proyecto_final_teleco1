package modelo

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"proyecto_teleco/utilidades"
)

type Tipo_ide string

const (
	Cedula              Tipo_ide = "CC"
	Extrangeria         Tipo_ide = "CE"
	Documento_identidad Tipo_ide = "TI"
)

type Usuario struct {
	ID         uint           `json:"ID,omitempty" gorm:"primaryKey"`
	Nombre     string         `json:"Nombre"`
	Apellido   string         `json:"Apellido"`
	Telefono   int            `json:"Telefono"`
	Tipo_id    Tipo_ide       `json:"Tipo_id" gorm:"index:idx_member,unique"`
	Num_ide    uint           `json:"Num_ide" gorm:"index:idx_member,unique"`
	Admin      bool           `json:"admin,omitempty" `
	Texto      string         `json:"Texto"`
	Correo     string         `json:"Correo"`
	Clave1     string         `json:"Clave"`
	Id_usuario []JWT_database `json:"-" gorm:"foreignKey:Token"`
}

func (U *Usuario) Is_admin() bool {
	return U.Admin
}
func (U *Usuario) Set_admin() {
	U.Admin = true
}

func (U *Usuario) Validar_llave(clave string) error {
	var nueva_clave string = base64.StdEncoding.EncodeToString(
		[]byte(utilidades.GenerarSHA256_with_32bits(clave)))

	if nueva_clave == U.Clave1 {
		return nil
	}
	return errors.New("error,llave incorrecta ")
}

func Validaciones(U *Usuario) error {
	log.Println(U.Texto)
	if U.Texto == "" {
		return errors.New("error Texto  Vacio")
	}
	if U.Clave1 == "" {
		return errors.New("error Clave  Vacio")
	}
	if len(U.Clave1) > 32 {
		return errors.New("error Clave  con formatoi incorrecto ,por favor dijitar una clave con numeros y letras menor a 32 caracteres")
	}
	if !utilidades.Validar_correo(U.Correo) {
		return errors.New("error Correo Incorrecto")
	}
	if !utilidades.Validar_telefono(fmt.Sprint(U.Telefono)) {
		return errors.New("error: Telefono Incorrecto ,porfavor un numero de exactamente 10 digitos")
	}
	if !utilidades.Validar_clave_signos(U.Clave1) {
		return errors.New("error Clave  con formatoi incorrecto ,por favor dijitar una clave con numeros y letras menor a 32 caracteres")
	}

	return nil

}

func (U *Usuario) Validar_usuario() error {

	if errv := Validaciones(U); errv != nil {
		return errv
	}
	clave2 := utilidades.GenerarSHA256_with_32bits(U.Clave1)
	enc, erra := utilidades.EncryptAES(clave2, U.Texto)
	if erra != nil {
		return errors.New("error, texto no pudo ser encriptado")
	}
	U.Texto = enc
	U.Clave1 = base64.StdEncoding.EncodeToString([]byte(clave2))
	return nil
}

func (U *Usuario) CheckUpdate_usuario(Un *Usuario) error {
	if errv := Validaciones(Un); errv != nil {
		return errv
	}

	ct, _ := base64.StdEncoding.DecodeString(U.Clave1)
	tt, _ := base64.StdEncoding.DecodeString(U.Texto)

	if U.Clave1 != Un.Clave1 {

		txt, err2 := utilidades.DecryptAES(string(ct), string(tt))
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

	return errors.New("error , porfavor enviar los datos correctamente ,")

}
