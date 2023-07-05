package utilidades

import (
	"regexp"
)

func Validar_correo(text string) bool {
	var correo_validar string = "[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}"
	var validador *regexp.Regexp = regexp.MustCompile(correo_validar)
	return validador.MatchString(text)
}
func Validar_clave_signos(text string) bool {
	var validar string = "(\\d+)|[a-zA-Z]+|\\%|\\@"
	var validador *regexp.Regexp = regexp.MustCompile(validar)
	var total string = validador.ReplaceAllLiteralString(text, "")
	return len(total) == 0
}
func Validar_telefono(text string) bool {
	var validar string = "\\d{10}"
	var validador *regexp.Regexp = regexp.MustCompile(validar)
	return validador.MatchString(text)

}
