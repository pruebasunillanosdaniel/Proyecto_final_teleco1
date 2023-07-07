package main

import (
	"log"
	"proyecto_teleco/controlador"
	"proyecto_teleco/routes"
)

func main() {
	log.Println("v ...")
	err := controlador.Create_admin()
	if err != nil {
		log.Println(err)
	}
	routes.Crear_servidor()

}
