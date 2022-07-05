package facote

import (
	"fmt"
	"testing"
)

func TestCreateHelper(t *testing.T) {
	ch := CreateHelper(4, GetOption())
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 3, 2, 1}
	println(fmt.Sprintf("初始数据 %d", data))
	encrypt := ch.Encrypt(data)
	println(fmt.Sprintf("密数据 %d", encrypt))
	decrypt := ch.Decrypt(encrypt)
	println(fmt.Sprintf("解密数据 %d", decrypt))
}

func TestHelperMoreToken(t *testing.T) {
	ch := CreateHelper(10, GetOption())
	for i := 0; i < 100; i++ {
		uca := 1000000 + i
		encrypt, err := ch.CreateToken(uca)
		if err != nil {
			println(fmt.Sprintf("加密 %s", err))
			return
		}
		ucb, err := ch.CheckToken(encrypt, true)
		if err != nil {
			println(fmt.Sprintf("解密 %s", err))
			return
		}
		println(fmt.Sprintf("Uca: %d, Ucb: %d, Encrypt: %s", uca, ucb, encrypt))
	}
}
