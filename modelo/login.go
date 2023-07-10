package modelo

import (
	"errors"
	"fmt"
	"proyecto_teleco/utilidades"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT_DATA struct {
	jwt.RegisteredClaims
	Fecha_inicio time.Time `json:"Fecha_inicio"`
	ID           uint      `json:"ID"`
}

type JWT_database struct {
	ID        uint `gorm:"primaryKey"`
	Insercion time.Time
	Datos     string
	Token     uint
}

func Create_Jwt_database(u Usuario) (JWT_database, error) {
	var JJ JWT_database
	var err error

	DD := JWT_DATA{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Minute)),
			Subject:   "Usuario" + fmt.Sprint(u.ID),
		},
		Fecha_inicio: time.Now(),
		ID:           u.ID,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &DD)

	JJ.Datos, err = t.SignedString([]byte(utilidades.Secret_jwt))
	JJ.Insercion = time.Now()
	return JJ, err
}

// da el usuario al que referecia JWT TOKEN
func Get_JWT_ID(clave string) (uint, error) {
	var DD JWT_DATA
	t, err := jwt.ParseWithClaims(clave, &DD, func(t *jwt.Token) (interface{}, error) {
		return []byte(utilidades.Secret_jwt), nil
	})
	if !t.Valid {
		return 0, errors.New("mal token ,no valido ,intente genera token nuevo")
	}
	return DD.ID, err
}

func (DD *JWT_database) Is_valid() bool {

	var salida JWT_DATA
	t, _ := jwt.ParseWithClaims(DD.Datos, salida, func(t *jwt.Token) (interface{}, error) {
		return []byte(utilidades.Secret_jwt), nil
	})
	return t.Valid
}
