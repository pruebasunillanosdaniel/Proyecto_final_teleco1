package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var database_Azure *gorm.DB = nil

func connect() (*gorm.DB, error) {

	var host string = os.Getenv("Host_sqlserver")
	port, err := strconv.Atoi(os.Getenv("Port_sqlserver"))
	if err != nil {
		return nil, errors.New("error: puerto incorrecto")
	}
	var database string = os.Getenv("Database_sqlserver")
	var user string = os.Getenv("User_sqlserver")
	var password string = os.Getenv("Password_sqlserver")
	fmt.Println(host, port, database, user, password)
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d ;database=%s;",
		host, user, password, port, database)
	fmt.Println(dsn)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	return db, err
}

func Database() *gorm.DB {

	if database_Azure == nil {
		db, err := connect()
		if err != nil {
			return nil
		}
		database_Azure = db
		return database_Azure
	}
	return database_Azure
}