package main

import (
	"log"
	"proyecto_teleco/database"
)

func main() {
	log.SetFlags(log.Llongfile)
	//log.Println(controlador.Enviar("solrevisdor143@gmail.com", "envio", "hola"))
	log.Println(database.Database())
}
