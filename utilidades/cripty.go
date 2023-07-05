package utilidades

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"time"
)

// tomado de https://www.golangprograms.com/advanced-encryption-standard.html

func EncryptAES(key string, texto string) (string, error) {

	if len(key) != 32 {
		if len(key) < 32 {
			key = strings.Trim(strings.ReplaceAll(fmt.Sprint(make([]int, 32-len(key))), " ", ""), "[]") + key
		} else {
			return "", errors.New("Clave superior a lo permitido por favor una clave menor a 32 caracteres")
		}
	}
	var texto_stream []byte = []byte(texto)
	texto_stream = []byte(texto)
	blockCipher, _ := aes.NewCipher([]byte(key))
	stream := cipher.NewCTR(blockCipher, []byte(Clave_texto16))
	stream.XORKeyStream(texto_stream, texto_stream)

	return string(texto_stream), nil

}

func DecryptAES(key string, ct string) (string, error) {
	if len(key) != 32 {
		if len(key) < 32 {
			key = strings.Trim(strings.ReplaceAll(fmt.Sprint(make([]int, 32-len(key))), " ", ""), "[]") + key
		} else {
			return "", errors.New("Clave superior a lo permitido por favor una clave menor a 32 caracteres")
		}
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	var CPtext []byte = []byte(ct)
	mode := cipher.NewCTR(block, []byte(Clave_texto16))
	mode.XORKeyStream(CPtext, CPtext)
	return string(CPtext), nil
}

func CompareAES(key string, text_compare string, key_compare string) (bool, error) {

	encriptado1, err := DecryptAES(key, text_compare)
	if err != nil {
		return false, err
	}
	encriptado2, err2 := DecryptAES(key_compare, text_compare)
	if err2 != nil {
		return false, err2
	}

	if encriptado2 == encriptado1 {
		return true, nil
	}
	return false, errors.New("Llaves no coiciden")

}

func GenerarSHA254(texto string) string {
	conv := sha256.Sum256([]byte(texto))
	return string(conv[:])
}

func GenerarShawithTime() string {
	tiempo := time.Now().String()
	return GenerarSHA254(tiempo)
}
