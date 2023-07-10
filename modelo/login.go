package modelo

import (
	"errors"
	"fmt"
	"proyecto_teleco/utilidades"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   fmt.Sprint(u.ID),
			ID:        fmt.Sprint(u.ID),
			Audience:  []string{"teleinformatica"},
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
func Get_JWT_ID_usuario(clave string) (uint, error) {
	var DD JWT_DATA
	t, err := jwt.ParseWithClaims(clave, &DD, func(t *jwt.Token) (interface{}, error) {
		return []byte(utilidades.Secret_jwt), nil
	})

	if err != nil {
		return 0, err
	}

	if t.Valid {
		return DD.ID, nil
	}
	return 0, errors.New("token no valido")
}

func Get_JWT(clave string) (JWT_DATA, error) {
	var salida JWT_DATA
	t, _ := jwt.ParseWithClaims(clave, salida, func(t *jwt.Token) (interface{}, error) {
		return []byte(utilidades.Secret_jwt), nil
	})

	if k, ok := t.Claims.(JWT_DATA); ok && t.Valid {
		return k, nil
	}
	return JWT_DATA{}, nil
}

func (DD *JWT_database) Is_valid() bool {

	var salida jwt.MapClaims
	t, err := jwt.ParseWithClaims(DD.Datos, salida, func(t *jwt.Token) (interface{}, error) {
		return []byte(utilidades.Secret_jwt), nil
	})

	if claims, ok := t.Claims.(*JWT_DATA); ok && t.Valid {
		fmt.Printf("%v ", claims.RegisteredClaims.Issuer)
		return true
	}

	fmt.Println("error en token", err)
	return false

}
