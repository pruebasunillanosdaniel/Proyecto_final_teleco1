package main

import (
	"log"
	"proyecto_teleco/controlador"
	"proyecto_teleco/routes"
)

func main() {
	log.Println("v ...")
	err := controlador.Create_init()
	if err != nil {
		log.Println(err)
	}
	routes.Crear_servidor()

}
