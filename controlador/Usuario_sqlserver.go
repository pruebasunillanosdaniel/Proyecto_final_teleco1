package controlador

import (
	"errors"
	"proyecto_teleco/database"
	"proyecto_teleco/modelo"
)

func Crear_usuario(U *modelo.Usuario) error {

	db, _ := database.Database()
	if db != nil {
		return errors.New("errores creando Usuario")
	}
	err := db.Create(U).Error
	return err
}

func Update_usuario(U *modelo.Usuario) error {

	db, _ := database.Database()
	if db != nil {
		return errors.New("errores creando Usuario")
	}
	err := db.Save(U).Error
	return err

}

func Read_usuario(id uint) (modelo.Usuario, error) {

	db, _ := database.Database()
	var U modelo.Usuario
	if db != nil {
		return U, errors.New("errores creando Usuario")
	}

	err := db.Where("ID=?", id).First(&U).Error
	return U, err

}

func Delete_usuario(id uint) (modelo.Usuario, error) {

	db, _ := database.Database()
	var U modelo.Usuario
	if db != nil {
		return U, errors.New("errores creando Usuario")
	}

	err := db.Where("ID=?", id).Delete(&U).Error
	return U, err

}

func List_all_user() ([]modelo.Usuario, error) {

	db, _ := database.Database()
	var U []modelo.Usuario
	if db != nil {
		return nil, errors.New("errores creando Usuario")
	}
	err := db.Find(&U).Error
	return U, err
}

func Get_User_by_unique(tipo_ide modelo.Tipo_ide, identificacion uint) (modelo.Usuario, error) {

	db, _ := database.Database()
	var u modelo.Usuario
	err := db.Model(&modelo.Usuario{}).
		Where(&modelo.Usuario{Num_ide: identificacion, Tipo_id: tipo_ide}).
		First(&u).Error

	if err != nil {
		return modelo.Usuario{}, errors.New("NO existe el usuario")
	}

	return u, nil
}
