package controlador

import (
	"errors"
	"proyecto_teleco/database"
	"proyecto_teleco/modelo"

	"gorm.io/gorm"
)

func Crear_JWT(U *modelo.JWT_database) error {

	db, err := database.Database()
	if err != nil {
		return errors.New("errores creando Usuario, fallo conexion DB ")
	}
	err = db.Table("Jwt_databases").Create(U).Error
	return err
}

func Crear_exist_JWT(U *modelo.JWT_database) error {

	db, err := database.Database()
	if err != nil {
		return errors.New("errores creando Usuario, fallo conexion DB ")
	}
	err = db.Table("Jwt_databases").FirstOrCreate(U).Error
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
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return modelo.JWT_database{}, errors.New("NO existe el Token ,porfavor genere uno")
	}
	return U, err

}

func Read_JWT_usuario(id_usuario uint) (modelo.JWT_database, error) {
	var err error
	db, err := database.Database()
	if err != nil {
		return modelo.JWT_database{}, errors.New("errores leyendo Usuario, fallo conexion DB ")
	}

	var Ua modelo.JWT_database = modelo.JWT_database{}

	err = db.Model(&modelo.JWT_database{}).Where("token=?", id_usuario).First(&Ua).Error
	if err != nil {
		return modelo.JWT_database{}, errors.New("errores no existe token ")
	}
	return Ua, err

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

	err = db.Where("token=?", id_usuario).Delete(&U).Error
	return err

}
