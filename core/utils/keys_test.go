package utils

import (
	"testing"
)

func TestCrypted(t *testing.T) {
	cryptedStr := Crypted("Password1")
	txt := DeCryptedStr(cryptedStr)
	t.Log(txt)
}
