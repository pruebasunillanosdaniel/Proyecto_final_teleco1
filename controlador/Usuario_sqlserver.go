package controlador

import (
	"errors"
	"proyecto_teleco/database"
	"proyecto_teleco/modelo"
)

func Crear_usuario(U *modelo.Usuario) error {

	db, err := database.Database()
	if err != nil {
		return errors.New("errores creando Usuario, fallo conexion DB ")
	}
	err = db.Table("Usuarios").Create(U).Error
	return err
}

func Update_usuario(U *modelo.Usuario) error {

	db, err := database.Database()
	if err != nil {
		return errors.New("errores creando Usuario, fallo conexion DB ")
	}
	err = db.Save(U).Error
	return err

}

func Read_usuario(id uint) (modelo.Usuario, error) {

	db, err := database.Database()
	var U modelo.Usuario
	if err != nil {
		return modelo.Usuario{}, errors.New("errores creando Usuario, fallo conexion DB ")
	}

	err = db.Where("ID=?", id).First(&U).Error
	return U, err

}

func Delete_usuario(id uint) error {

	db, err := database.Database()
	var U modelo.Usuario
	if err != nil {
		return errors.New("errores creando Usuario, fallo conexion DB ")
	}

	err = db.Where("ID=?", id).Delete(&U).Error
	return err

}

func List_all_user() ([]modelo.Usuario, error) {

	db, err := database.Database()
	var U []modelo.Usuario
	if err != nil {
		return nil, errors.New("errores listando Usuario,no existen")
	}
	err = db.Find(&U).Error
	return U, err
}

func Get_User_by_unique(tipo_ide modelo.Tipo_ide, identificacion uint) (modelo.Usuario, error) {

	db, err := database.Database()
	var u modelo.Usuario
	if err != nil {
		return modelo.Usuario{}, errors.New("errores DB conexion")
	}
	err = db.Model(&modelo.Usuario{}).
		Where(&modelo.Usuario{Num_ide: identificacion, Tipo_id: tipo_ide}).
		First(&u).Error

	if err != nil {
		return modelo.Usuario{}, errors.New("NO existe el usuario")
	}

	return u, nil
}

func Create_admin() error {

	db, err := database.Database()
	var u modelo.Usuario = modelo.Usuario{
		ID:       1,
		Nombre:   "admin",
		Apellido: "admin",
		Correo:   "solrevisdor143@gmail.com",
		Telefono: 0000000000,
		Num_ide:  0,
		Tipo_id:  "CC",
		Clave1:   "8C6976E5B5410415BDE908BD4DEE15DFB167A9C873FC4BB8A81F6F2AB448A918",
		Texto:    "",
	}

	if err != nil {
		return err
	}
	err = db.Model(&modelo.Usuario{}).FirstOrCreate(&u).Error
	return err
}
