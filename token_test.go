package facote

import (
	"fmt"
	"testing"
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
	uc, err := CheckToken("dGBoD10XAjMdSl0aJxhYKxdaShBRAxdvWjIPWjxAcBITWSwAVxQYBU1OcEUXEkwiShZdeUlYBkl5SnwPA1oRWnQDF1EAQVgOF2dQQWA9DgtdWTslBRZMDgtvcEBJEm8BSlYic2UDBnhlSjg2A0EhSkNXUAwOUDFRUAxTSgYNHA==")
	if err != nil {
		println(fmt.Sprintf("%s", err))
		return
	}
	println(fmt.Sprintf("Successful: uc: %d", uc))
}

func TestMoreCheckTokenShow(t *testing.T) {
	for i := 0; i < 100; i++ {
		//time.Sleep(10 * time.Millisecond)
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

func TestMoreCheckTokenTime(t *testing.T) {
	for i := 0; i < 1000; i++ {
		uca := 1000000 + i
		encrypt, err := CreateToken(uca, int64(0))
		if err != nil {
			println(fmt.Sprintf("加密 %s", err))
			return
		}
		_, err = CheckToken(encrypt)
		if err != nil {
			println(fmt.Sprintf("解密 %s", err))
			return
		}
		//println(fmt.Sprintf("Uca: %d, Ucb: %d, Encrypt: %s", uca, ucb, encrypt))
	}
}

// 获取乱序 Key
func TestKeyBytes(t *testing.T) {
	bytes := make([]byte, 0)
	for i := 0; i < 256; i++ {
		bytes = append(bytes, byte(i))
	}
	// 直接展示
	println(fmt.Sprintf("%d", bytes))
	ShuffleBytes(bytes)
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

// 获取非控制字符的乱序 Key 33 ... 91 93...126 乱序，总长度 128
func TestJsonKeyBytes(t *testing.T) {
	bytes := make([]byte, 0)
	for i := 33; i < 92; i++ {
		bytes = append(bytes, byte(i))
	}
	for i := 93; i < 127; i++ {
		bytes = append(bytes, byte(i))
	}
	println(fmt.Sprintf("需求长度：%d", len(bytes)))
	println(fmt.Sprintf("需求字符：%d", bytes))

	ShuffleBytes(bytes)

	x := bytes[:59]
	println(fmt.Sprintf("第一需求：%d", x))
	y := bytes[59:]
	println(fmt.Sprintf("第二需求：%d", y))

	data := make([]byte, 0)
	for i := 0; i < 33; i++ {
		data = append(data, byte(i))
	}
	println(fmt.Sprintf("头部空白：%d", data))
	data = append(data, x...)
	println(fmt.Sprintf("拼接一次：%d", data))
	data = append(data, 92)
	println(fmt.Sprintf("拼接两次：%d", data))
	data = append(data, y...)
	println(fmt.Sprintf("拼接三次：%d", data))
	data = append(data, 127)
	println(fmt.Sprintf("拼接两次：%d", data))

	// 用逗号隔开展示
	byteStr := "需求展示：["
	for i, v := range data {
		if i == 0 {
			byteStr = fmt.Sprintf("%s%d", byteStr, v)
		} else {
			byteStr = fmt.Sprintf("%s, %d", byteStr, v)
		}
	}
	byteStr = fmt.Sprintf("%s]  ", byteStr)
	println(byteStr)

	// 顺序查看
	for i, v := range data {
		println(fmt.Sprintf("%d - %d", i, v))
	}

}

func TestBytes128(t *testing.T) {
	bytes := make([]byte, 0)
	for i := 0; i < 128; i++ {
		bytes = append(bytes, byte(i))
	}

	ShuffleBytes(bytes)

	// 用逗号隔开展示
	byteStr := "需求展示：["
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
