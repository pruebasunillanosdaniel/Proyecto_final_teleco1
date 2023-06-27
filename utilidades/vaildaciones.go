package utilidades

import (
	"regexp"
)

func Validar_correo(text string) bool {
	var correo_validar string = "[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}"
	var validador *regexp.Regexp = regexp.MustCompile(correo_validar)
	return validador.MatchString(text)
}
