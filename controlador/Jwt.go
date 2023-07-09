package controlador

import (
	"errors"
	"proyecto_teleco/database"
	"proyecto_teleco/modelo"
)

func Crear_JWT(U *modelo.JWT_database) error {

	db, err := database.Database()
	if err != nil {
		return errors.New("errores creando Usuario, fallo conexion DB ")
	}
	err = db.Table("Jwt_databases").Create(U).Error
	return err
}

func Update_JWT(U *modelo.JWT_database) error {

	db, err := database.Database()
	if err != nil {
		return errors.New("errores creando Usuario, fallo conexion DB ")
	}
	err = db.Save(U).Error
	return err

}

func Read_JWT(id uint) (modelo.JWT_database, error) {

	db, err := database.Database()
	var U modelo.JWT_database
	if err != nil {
		return modelo.JWT_database{}, errors.New("errores creando Usuario, fallo conexion DB ")
	}

	err = db.Where("id=?", id).First(&U).Error
	return U, err

}

func Read_JWT_token(token string) (modelo.JWT_database, error) {

	db, err := database.Database()
	var U modelo.JWT_database
	if err != nil {
		return modelo.JWT_database{}, errors.New("errores creando Usuario, fallo conexion DB ")
	}
	err = db.Where("datos=?", token).First(&U).Error
	return U, err

}

func Read_JWT_usuario(id_usuario uint) (modelo.JWT_database, error) {

	db, err := database.Database()
	var U modelo.JWT_database
	if err != nil {
		return modelo.JWT_database{}, errors.New("errores leyendo Usuario, fallo conexion DB ")
	}
	err = db.Where("usuario=?", id_usuario).First(&U).Error
	return U, err

}

func Delete_JWT(id uint) error {

	db, err := database.Database()
	var U modelo.JWT_database
	if err != nil {
		return errors.New("errores eliminando Usuario, fallo conexion DB ")
	}

	err = db.Where("id=?", id).Delete(&U).Error
	return err

}

func Delete_JWT_usuario(id_usuario uint) error {

	db, err := database.Database()
	var U modelo.JWT_database
	if err != nil {
		return errors.New("errores eliminando Usuario, fallo conexion DB ")
	}

	err = db.Where("usuario=?", id_usuario).Delete(&U).Error
	return err

}
