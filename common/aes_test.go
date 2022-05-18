package common

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	key := []byte{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7}
	data := []byte{0, 1, 2, 3, 5, 23, 20, 11}
	bytes, err := Encrypt(data, key)
	if err != nil {
		println(fmt.Sprintf("%s", err))
		return
	}
	println(fmt.Sprintf("%d", bytes))
}
