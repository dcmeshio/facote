package facote

import (
	"fmt"
	"testing"
)

func TestCreateHelper(t *testing.T) {
	var key = []byte{91, 75, 80, 77, 31, 7, 23, 62, 64, 51, 112, 103, 48, 46, 13, 52, 122, 15, 98, 59, 65, 54, 95, 35, 101, 1, 71, 82, 26, 18, 32, 92, 100, 42, 2, 36, 43, 102, 45, 27, 68, 41, 39, 84, 55, 0, 116, 61, 4, 108, 30, 50, 11, 93, 85, 19, 34, 107, 89, 113, 114, 47, 5, 72, 117, 104, 44, 121, 14, 111, 6, 67, 20, 69, 17, 119, 115, 81, 53, 97, 76, 126, 125, 9, 33, 96, 83, 88, 105, 123, 86, 16, 49, 124, 127, 79, 94, 87, 60, 99, 25, 109, 90, 38, 56, 74, 118, 21, 28, 63, 22, 12, 3, 24, 70, 110, 37, 10, 40, 57, 106, 73, 8, 78, 58, 120, 29, 66}
	CreateHelper(2, "73e4f31e-b5ce-4c44-8078-e87bdf267116", key)
}

func TestCheck(t *testing.T) {
	var key = []byte{91, 75, 80, 77, 31, 7, 23, 62, 64, 51, 112, 103, 48, 46, 13, 52, 122, 15, 98, 59, 65, 54, 95, 35, 101, 1, 71, 82, 26, 18, 32, 92, 100, 42, 2, 36, 43, 102, 45, 27, 68, 41, 39, 84, 55, 0, 116, 61, 4, 108, 30, 50, 11, 93, 85, 19, 34, 107, 89, 113, 114, 47, 5, 72, 117, 104, 44, 121, 14, 111, 6, 67, 20, 69, 17, 119, 115, 81, 53, 97, 76, 126, 125, 9, 33, 96, 83, 88, 105, 123, 86, 16, 49, 124, 127, 79, 94, 87, 60, 99, 25, 109, 90, 38, 56, 74, 118, 21, 28, 63, 22, 12, 3, 24, 70, 110, 37, 10, 40, 57, 106, 73, 8, 78, 58, 120, 29, 66}
	ch := CreateHelper(2, "73e4f31e-b5ce-4c44-8078-e87bdf267116", key)
	token := "bXQ9MFhqZVoMWyF5NylrXR4lTjgoAglMFVF1bVdkSXZ1AmcSWSAcAgITBjtYPwESGSF9IUssbQAtUzdrY0MvOEdlBhwQFCxVJVQuDnEUe3pSCRAxblslCxg/EWwQFTlkOUkyJm4tbyUZfihXJnosLDNQd2RCDl08DDsKVjkcDnYtaRNdOFgWBj99LBojOERPEWEsJixBaRxcdgM5ez1RYWtrIWALdjs7bjUGFVZOd30UU2FxBFtqHFt1ex5OMhdSVncqexRLYx4mSUdqERREPElcVxk+Shd4BDVPGV4pWnhBOXgUflN5dmwhIT4CaRE6XTEnIkR+c2tednwAaxgiajcFaH5uOWgBah01VhUfb3tHCFxfUwVDfmZ1aVJ4BmsnfhpobV9ib2FgDlgTRRV9VyIWVnlCFVo+LAllFSB8aGN/FQpgRX8mBi02CXNPQ1MwazkjOR0IBnUmfRFFIgQmKF5BfnNOWFR+RUESMiRPYARyIR9PUQhgOQUIWFtAPy5UBGY1VEYeVh0JfkELFVImXWg7WlcSIWsYBlQUQ2hEaUJjBEEHJX5gCUhYQxEAaDwtA1NNEQ8XPAtuSiBVGR1YVXQHDkcTMi4+U2RUUwhITFMoaQV7EVkrBDFPaGYmLHB7JSplMmpWBlISSR4XAxJHWC5Tc2wqLEM0OXETc2dgPBU3DmhYanAGTCBIVxk0IR1tfWsJRjYQfSZUHhJzOFMUb2hvCQA3JlxaYytZLyN7AmRxNyU/Ag0oAxoZbn8AAms9WWhqAk14ExU2MkN1bT4YCz17WhJTMjVFbB0kbWR7ACkzPBk1XVYFYwY9bTw6AFpzC14HYwkMCyokCy8IAA4wIgEJBAEAE3E9IngyACRwbRpYImNAEwxbPDMdGSZqWkFLHkoxVTUkE2YkbAs3bB9xVVcBAmsVN0kaAk0YJS5UbjYvAh50WTQ3bB5WVWpFXQhXEy8QbB8yBBpzMhY4IgFlBFoXVXkHeA=="
	ck, err := ch.CheckToken(token, false)
	if err != nil {
		println(fmt.Sprintf("解密失败，%s", err))
		return
	}
	println(fmt.Sprintf("结果：%d", ck.Key))
}

func TestCryptoHelper(t *testing.T) {
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
