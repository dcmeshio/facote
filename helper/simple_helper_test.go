package helper

import (
	"fmt"
	"testing"
)

// 客户端 通过 sch 创建私钥
// 通过网络将 私钥口令传给 服务端
// 服务端 通过私钥口令 获取私钥
// 服务端 数据私钥加密数据

func TestCheck(t *testing.T) {
	sch := sch()
	// 客户端 通过 sch 创建私钥
	ck, err := sch.createKey()
	if err != nil {
		println(fmt.Sprintf("Error: %s", err))
		return
	}
	println(fmt.Sprintf("Token: %s", ck.Token))
	println(fmt.Sprintf("CK key: [% x]", ck.Key))

	// 服务端 通过私钥口令 获取私钥
	dck, err := sch.CheckToken(ck.Token)
	if err != nil {
		println(fmt.Sprintf("Error: %s", err))
		return
	}
	println(fmt.Sprintf("Token: %s", dck.Token))
	println(fmt.Sprintf("CK key: [% x]", dck.Key))

	// 创建原始数据
	data := []byte("abc")
	println(fmt.Sprintf("初始数据 %d", data))

	// 原始数据 私钥加密
	data = dck.Encrypt(data)
	// 原始数据 公钥加密
	data = sch.Encrypt(data)
	println(fmt.Sprintf("加密数据 %d", data))

	// 加密数据 公钥解密
	data = sch.Decrypt(data)
	// 加密数据 私钥解密
	data = ck.Decrypt(data)
	println(fmt.Sprintf("解密数据 %s", string(data)))

}

func TestRes(t *testing.T) {
	sch := sch()
	token := "TgIVbUkCWQJtFmNKWkoZFmgIBnZXYGMDTGpFLAkVXUNpaDgLfkluHEVJLBwZFAR3FDgsUzwJOQwJCFMGaSYLOCw5BGk1FmAlWFNRCFc4VkwOaEURViYleRQhEW9zCCx9LFgmGFosJkMRDmxdDmwsSiEhagpzOCVJVwgVYCwOJTxgOGtgET8GaGMyImhDYG8EYQ4DHhRDOTIJU2MhVhQlJkM/a3tvU1ZqaWBqV0wcIlMJFGgvAjcCA24CWQJoPwwEBnY4GGAIfh5obyJWNQ57WEN9e1g1JhxWaCYVHjVoFXlob28LUTk5eWBsUVECeA=="
	// 服务端 通过私钥口令 获取私钥
	ck, err := sch.CheckToken(token)
	if err != nil {
		println(fmt.Sprintf("Error: %s", err))
		return
	}
	println(fmt.Sprintf("Token: %s", ck.Token))
	println(fmt.Sprintf("CK key: %d", ck.Key))
	println(fmt.Sprintf("CK unkey: %d", ck.Unkey))

	// 创建原始数据

	data := []byte{1, 2, 3}
	println(fmt.Sprintf("初始数据 %d", data))

	// 原始数据 私钥加密
	data = ck.Encrypt(data)
	// 原始数据 公钥加密
	data = sch.Encrypt(data)
	println(fmt.Sprintf("加密数据 %d", data))

	// 加密数据 公钥解密
	data = sch.Decrypt(data)
	// 加密数据 私钥解密
	data = ck.Decrypt(data)
	println(fmt.Sprintf("解密数据 %d", data))
}

func sch() *SimpleCryptoHelper {
	var key = []byte{91, 75, 80, 77, 31, 7, 23, 62, 64, 51, 112, 103, 48, 46, 13, 52, 122, 15, 98, 59, 65, 54, 95, 35, 101, 1, 71, 82, 26, 18, 32, 92, 100, 42, 2, 36, 43, 102, 45, 27, 68, 41, 39, 84, 55, 0, 116, 61, 4, 108, 30, 50, 11, 93, 85, 19, 34, 107, 89, 113, 114, 47, 5, 72, 117, 104, 44, 121, 14, 111, 6, 67, 20, 69, 17, 119, 115, 81, 53, 97, 76, 126, 125, 9, 33, 96, 83, 88, 105, 123, 86, 16, 49, 124, 127, 79, 94, 87, 60, 99, 25, 109, 90, 38, 56, 74, 118, 21, 28, 63, 22, 12, 3, 24, 70, 110, 37, 10, 40, 57, 106, 73, 8, 78, 58, 120, 29, 66}
	return CreateHelper("73e4f31e-b5ce-4c44-8078-e87bdf267116", key)
}
