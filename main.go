package main

import (
	"log"
	"proyecto_teleco/controlador"
	"proyecto_teleco/database"
	"proyecto_teleco/modelo"
	"proyecto_teleco/routes"
)

func main() {
	db, _ := database.Database()
	db.AutoMigrate(&modelo.Usuario{})
	var U modelo.Usuario = modelo.Usuario{
		Nombre:  "admin",
		Clave1:  "admin",
		Tipo_id: "CC",
		Num_ide: 000000,
	}
	err := controlador.Crear_usuario(&U)
	if err != nil {
		log.Println(err)
	}

	routes.Crear_servidor()
	for {

	}
}
