package controlador

import (
	"errors"
	"log"
	"proyecto_teleco/database"
	"proyecto_teleco/modelo"

	"gorm.io/gorm"
)

func Crear_usuario(U *modelo.Usuario) error {

	db, err := database.Database()
	if err != nil {
		return errors.New("errores creando Usuario, fallo conexion DB ")
	}
	err = db.Model(&modelo.Usuario{}).Create(U).Error
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

	err = db.Where("id=?", id).First(&U).Error
	return U, err

}

func Delete_usuario(id uint) error {

	db, err := database.Database()
	var U modelo.Usuario
	if err != nil {
		return errors.New("errores eliminando Usuario, fallo conexion DB ")
	}

	err = db.Where("id=?", id).Delete(&U).Error
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

func Get_User_by_ID(ID uint) (modelo.Usuario, error) {

	db, err := database.Database()
	var u modelo.Usuario
	if err != nil {
		return modelo.Usuario{}, errors.New("errores DB conexion")
	}
	err = db.Model(&modelo.Usuario{}).
		Where("id =? ", ID).
		First(&u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return modelo.Usuario{}, errors.New("no existe el usuario")
	} else if err != nil {
		return modelo.Usuario{}, err

	}

	return u, nil
}

func Create_init() error {

	db, err := database.Database()
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&modelo.Usuario{}, &modelo.JWT_database{})
	if err != nil {
		return err
	}
	var u modelo.Usuario = modelo.Usuario{
		ID:       1,
		Nombre:   "admin",
		Apellido: "admin",
		Correo:   "solrevisdor143@gmail.com",
		Telefono: 1234567890,
		Num_ide:  0,
		Tipo_id:  "CC",
		Clave1:   "admin",
		Texto:    "Hola mundo",
		Admin:    true,
	}

	err = u.Validar_usuario()
	if err != nil {
		return err
	}
	log.Println("paso usuario")
	u.Set_admin()
	if err != nil {
		return err
	}
	err = db.Model(&modelo.Usuario{}).FirstOrCreate(&u).Error
	if errors.Is(gorm.ErrDuplicatedKey, err) {
		return nil
	}
	return err
}
