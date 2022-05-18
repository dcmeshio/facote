package facote

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {

	x, err := CreateToken(100003, int64(0))
	if err != nil {
		println(fmt.Sprintf("%s", err))
		return
	}
	println(fmt.Sprintf("Encrypt: %s", x))

	// xYZA0L/bYC6CyDOZZF18vj026tFtTXXiQYHOd8EqfWs=

	//uc, qc, err := CheckToken("CQmzs8xf/hSSugETy7d5QUM2WEmI7k9TmUg7tW1EUFbVW6orJPiYPv7+XEIn/zU=")
	//if err != nil {
	//	println(fmt.Sprintf("error: %s", err))
	//	return
	//}
	//println(fmt.Sprintf("Successful: uc: %d, qc: %s", uc, qc))

}

func TestCheckToken(t *testing.T) {
	uc, err := CheckToken("yvq6/viC/j49rZVyRGnq4sxA708zcoRGID7wXmbV68r6dcr6vcr6qsr6b8r6vcr6Qsr6+cr6Ncr6Ncr6Ncr6Ncr6zcr6mMr6vcr6Zcr6Osr6vcr6Qsr6+cr6fcr6u8r6Wsr6F8r6P8r6Ncr6Wsr6u8r6fcr6mMr6vcr6lsr6Osr6vcr6Qsr6vcr6vcr6Gg==")
	if err != nil {
		println(fmt.Sprintf("%s", err))
		return
	}
	println(fmt.Sprintf("Successful: uc: %d", uc))
}

func TestMoreCheckToken(t *testing.T) {

	for i := 0; i < 1000; i++ {
		time.Sleep(1 * time.Second)
		uca := 1000000 + i
		encrypt, err := CreateToken(uca, int64(0))
		if err != nil {
			println(fmt.Sprintf("加密 %s", err))
			return
		}
		ucb, err := CheckToken(encrypt)
		if err != nil {
			println(fmt.Sprintf("解密 %s", err))
			return
		}
		println(fmt.Sprintf("Uca: %d, Ucb: %d, Encrypt: %s", uca, ucb, encrypt))
	}

}

// 获取乱序 Key
func TestKeyBytes(t *testing.T) {
	bytes := make([]byte, 0)
	for i := 0; i < 256; i++ {
		bytes = append(bytes, byte(i))
	}
	ShuffleBytes(bytes)
	// 直接展示
	println(fmt.Sprintf("%d", bytes))
	// 用逗号隔开展示
	byteStr := "["
	for i, v := range bytes {
		if i == 0 {
			byteStr = fmt.Sprintf("%s%d", byteStr, v)
		} else {
			byteStr = fmt.Sprintf("%s, %d", byteStr, v)
		}
	}
	byteStr = fmt.Sprintf("%s]  ", byteStr)
	println(byteStr)
}

func TestEncrypt(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 3, 2, 1}
	println(fmt.Sprintf("初始数据 %d", data))
	encrypt := Encrypt(data)
	println(fmt.Sprintf("密数据 %d", encrypt))
	decrypt := Decrypt(encrypt)
	println(fmt.Sprintf("解密数据 %d", decrypt))
	// 9, 12, 15, 18
}

func TestTwoBytes(t *testing.T) {
	bytes := twoBytes()
	println(fmt.Sprintf("%d", bytes))
}

func Test_tokenHeader(t *testing.T) {
	bytes := tokenHeader()
	println(fmt.Sprintf("%d", bytes))
}
