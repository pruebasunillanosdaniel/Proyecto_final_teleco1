package database

import (
	"fmt"
	"proyecto_teleco/utilidades"
	"strconv"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var database_Azure *gorm.DB = nil

func connect() (*gorm.DB, error) {

	var host string = utilidades.Azure_host
	port, _ := strconv.Atoi(utilidades.Azure_port)
	var database string = utilidades.Azure_db
	var user string = utilidades.Azure_user
	var password string = utilidades.Azure_password

	var dsn string = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d ;database=%s;",
		host, user, password, port, database)
	fmt.Println(dsn)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	//db.Callback().Create().Remove("mssql:set_identity_insert")
	return db, err
}

func Database() (*gorm.DB, error) {

	if database_Azure == nil {
		db, err := connect()
		if err != nil {
			return &gorm.DB{}, err
		}
		database_Azure = db
		return database_Azure, nil
	}
	return database_Azure, nil
}
