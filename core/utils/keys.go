package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
var keytext = "guozhiw12306eidavidguodramemaker"

// Crypted is encryption your string.
func Crypted(cryptedStr string) []byte {
	byteStr := []byte(cryptedStr)
	c, err := aes.NewCipher([]byte(keytext))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(keytext), err)
		os.Exit(-1)
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cryptedString := make([]byte, len(byteStr))
	cfb.XORKeyStream(cryptedString, byteStr)
	return cryptedString
}

// DeCryptedStr is decryption your string
func DeCryptedStr(deCryptedStr []byte) string {
	c, err := aes.NewCipher([]byte(keytext))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(keytext), err)
		os.Exit(-1)
	}
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	deCryptedString := make([]byte, len(deCryptedStr))
	cfbdec.XORKeyStream(deCryptedString, deCryptedStr)
	return string(deCryptedString)
}
