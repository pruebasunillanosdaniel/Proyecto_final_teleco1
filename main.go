package main

import (
	"log"
	"proyecto_teleco/controlador"
	"proyecto_teleco/routes"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Llongfile)
	err := controlador.Create_init()
	if err != nil {
		log.Fatal(err)
	}
	routes.Crear_servidor()

}
