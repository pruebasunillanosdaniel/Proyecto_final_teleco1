package utilidades

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

func EncryptAES(key []byte, plaintext string) (string, error) {

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out), nil
}

func DecryptAES(key []byte, ct string) (string, error) {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)
	s := string(pt[:])
	return s, nil
}

func CompareAES(key []byte, text_compare string, key_compare []byte) (bool, error) {

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
