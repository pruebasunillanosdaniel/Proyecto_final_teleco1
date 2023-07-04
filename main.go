package main

import (
	"log"
	"proyecto_teleco/utilidades"
)

func main() {
	log.SetFlags(log.Llongfile)
	log.Println("****************************************************++++++")
	a := "    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	v, _ := utilidades.EncryptAES(utilidades.GenerarSHA254("CLAVE"), a)
	log.Println(v)
	z, err1 := utilidades.DecryptAES(utilidades.GenerarSHA254("CLAVE"), v)
	if err1 != nil {
		log.Panic(err1)
	}
	log.Println(z)
}
