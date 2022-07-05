package facote

import (
	"fmt"
	"testing"
)

func TestCreateHelper(t *testing.T) {
	key := []byte{26, 99, 12, 64, 121, 7, 21, 41, 33, 124, 122, 62, 11, 53, 123, 117, 13, 18, 9, 39, 16, 83, 14, 66, 125, 110, 96, 17, 72, 82, 32, 25, 79, 127, 31, 120, 54, 76, 101, 20, 113, 112, 36, 23, 8, 58, 94, 126, 47, 80, 78, 22, 106, 65, 27, 108, 104, 75, 60, 102, 111, 103, 92, 50, 100, 95, 59, 88, 98, 29, 0, 70, 74, 34, 85, 51, 63, 46, 73, 37, 67, 30, 52, 118, 109, 38, 87, 1, 57, 81, 6, 115, 44, 4, 89, 49, 119, 77, 68, 35, 114, 105, 42, 40, 116, 71, 28, 43, 97, 93, 84, 19, 61, 107, 48, 91, 5, 10, 69, 24, 2, 45, 90, 56, 3, 15, 86, 55}
	ps := "abc123.."
	ch := CreateHelper(4, ps, key)
	cka, err := ch.CreateToken()
	if err != nil {
		println(fmt.Sprintf("加密失败，%s", err))
		return
	}
	println(fmt.Sprintf("初始秘钥 [% x]", cka.Key))
	println(fmt.Sprintf("反秘钥 [% x]", cka.Unkey))
	println(fmt.Sprintf("传输口令 %s", cka.Token))
	ckb, err := ch.CheckToken(cka.Token, false)
	if err != nil {
		println(fmt.Sprintf("解密失败，%s", err))
		return
	}
	println(fmt.Sprintf("传输秘钥 [% x]", ckb.Key))
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	println(fmt.Sprintf("初始数据 %d", data))
	encrypt := ckb.Encrypt(data)
	println(fmt.Sprintf("密数据 %d", encrypt))
	decrypt := cka.Decrypt(encrypt)
	println(fmt.Sprintf("解密数据 %d", decrypt))
}

func TestHelperMoreToken(t *testing.T) {
	key := []byte{26, 99, 12, 64, 121, 7, 21, 41, 33, 124, 122, 62, 11, 53, 123, 117, 13, 18, 9, 39, 16, 83, 14, 66, 125, 110, 96, 17, 72, 82, 32, 25, 79, 127, 31, 120, 54, 76, 101, 20, 113, 112, 36, 23, 8, 58, 94, 126, 47, 80, 78, 22, 106, 65, 27, 108, 104, 75, 60, 102, 111, 103, 92, 50, 100, 95, 59, 88, 98, 29, 0, 70, 74, 34, 85, 51, 63, 46, 73, 37, 67, 30, 52, 118, 109, 38, 87, 1, 57, 81, 6, 115, 44, 4, 89, 49, 119, 77, 68, 35, 114, 105, 42, 40, 116, 71, 28, 43, 97, 93, 84, 19, 61, 107, 48, 91, 5, 10, 69, 24, 2, 45, 90, 56, 3, 15, 86, 55}
	ps := "abc123.."
	ch := CreateHelper(2, ps, key)
	for i := 0; i < 100; i++ {
		ck, err := ch.CreateToken()
		if err != nil {
			println(fmt.Sprintf("加密 %s", err))
			return
		}
		println(fmt.Sprintf("加密 [% x]", ck.Key))
		println(fmt.Sprintf("口令 %s", ck.Token))
		println(fmt.Sprintf("长度 %d", len(ck.Token)))
		key, err := ch.CheckToken(ck.Token, false)
		if err != nil {
			println(fmt.Sprintf("解密 %s", err))
			return
		}
		println(fmt.Sprintf("解密 [% x]", key))
	}
}

func TestLength(t *testing.T) {
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJjb2RlIjoiNTk0NTYyNCIsInNpZ25pblR5cGUiOiJXZWJVc2VyIiwiZXhwIjoxNjU5NjAzMDk1LCJpYXQiOjE2NTY5MjQ2OTV9.YXocfP81P86kJ3kqw8VFlmiRU7IqvFRMk5DzLo1dSKp2aAVhSQ0S2UQ6Z-6NSbnhxaLT2rq5c80vliVWTA7neWDQa26tDGeWsIiSQSEzSc5PVqqZUDtDwITbzm63MmD_0OIutIAg9IsrZnUe5cKf-0l7JBrLPAzgirFKX510aZg"
	println(fmt.Sprintf("长度：%d", len(token)))
}
