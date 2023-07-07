package utilidades

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
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

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	if len(texto) < aes.BlockSize {

		texto = strings.Trim(strings.ReplaceAll(fmt.Sprint(make([]int, aes.BlockSize-len(texto))), " ", ""), "[]") + key
	}

	out := make([]byte, len(texto))

	c.Encrypt(out, []byte(texto))

	return hex.EncodeToString(out), nil

}

func DecryptAES(key string, ct string) (string, error) {
	log.Println(key)

	if len(key) != 32 {
		if len(key) < 32 {
			key = strings.Trim(strings.ReplaceAll(fmt.Sprint(make([]int, 32-len(key))), " ", ""), "[]") + key
		} else {
			return "", errors.New("Clave superior a lo permitido por favor una clave menor a 32 caracteres")
		}
	}
	log.Println(key)
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	return string(pt[:]), nil
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

func GenerarSHA256_with_32bits(texto string) string {
	if len(texto) < 32 {
		texto = strings.Trim(strings.ReplaceAll(fmt.Sprint(make([]int, 32-len(texto))), " ", ""), "[]") + texto
	}

	conv := sha256.Sum256([]byte(texto))
	return base64.StdEncoding.EncodeToString(conv[:])
}

func GenerarShawithTime() string {
	tiempo := time.Now().String()
	return GenerarSHA256_with_32bits(tiempo)
}
