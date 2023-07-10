package utilidades

import (
	"regexp"
)

var Correo *regexp.Regexp = regexp.MustCompile("[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}")
var Llaves *regexp.Regexp = regexp.MustCompile("(\\d+)|[a-zA-Z]+|\\%|\\@")
var Telefono *regexp.Regexp = regexp.MustCompile("\\d{10}")

func Validar_correo(text string) bool {

	return Correo.MatchString(text)
}
func Validar_clave_signos(text string) bool {
	var total string = Llaves.ReplaceAllLiteralString(text, "")
	return len(total) == 0
}
func Validar_telefono(text string) bool {
	return Telefono.MatchString(text)
}
