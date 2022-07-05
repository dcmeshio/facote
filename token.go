package facote

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Token struct {
	Uc int    `json:"uc"`
	Ts int64  `json:"ts"`
	Ps string `json:"ps"`
}

func CreateToken(uc int, td int64) (string, error) {
	opt := GetOption()
	// 创建结构体
	t := &Token{}
	t.Uc = uc
	t.Ps = opt.Ps
	timestamp := time.Now().Unix()
	t.Ts = timestamp + td
	// 转换为 Json 字符串
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	encryptData := Encrypt(data)
	return base64.StdEncoding.EncodeToString(encryptData), nil
}

func CheckToken(t string) (int, error) {
	opt := GetOption()
	data, err := base64.StdEncoding.DecodeString(t)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Base64 decode error: %s", err))
	}
	decryptData := Decrypt(data)
	var dt *Token
	err = json.Unmarshal(decryptData, &dt)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Json unmarshal error: %s", err))
	}

	if dt.Ps != opt.Ps {
		return 0, errors.New("check token ps error")
	}
	when := dt.Ts
	now := time.Now().Unix()
	if now-when > 30 || now-when < -10 {
		return 0, errors.New("check token timestamp error")
	}
	return dt.Uc, nil
}

func Encrypt(data []byte) []byte {
	bytes := make([]byte, 0)
	// 添加一个随机长度的头部，再每个正确字符前添加两个随机字符
	head := tokenHeader()
	bytes = append(bytes, head...)
	for i := 0; i < len(data); i++ {
		bytes = append(bytes, twoBytes()...)
		bytes = append(bytes, data[i])
	}
	// 进行乱序
	key := GetOption().Key
	for i := 0; i < len(bytes); i++ {
		index := bytes[i]
		bytes[i] = key[index]
	}
	return bytes
}

func Decrypt(data []byte) []byte {
	unkey := GetOption().Unkey
	for i := 0; i < len(data); i++ {
		index := data[i]
		data[i] = unkey[index]
	}
	bytes := make([]byte, 0)
	var count int
	for i, v := range data {
		if i == 3 {
			count = int(v)
		}
		if (i-count-3) > 0 && (i-count-3)%3 == 0 {
			bytes = append(bytes, data[i])
		}
	}
	return bytes
}

// 加入一个 0 ~ 30 位的头部
// [A, A, A, 3, C, C, C]
// 开头为三个随机数值 A
// 第四个值为 C 的数量
// C 为随机值
func tokenHeader() []byte {
	key := GetOption().Key
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, 0)
	for i := 0; i < 3; i++ {
		index := r.Intn(len(key))
		bytes = append(bytes, key[index])
	}
	count := r.Intn(30)
	bytes = append(bytes, byte(count))
	for i := 0; i < count; i++ {
		index := r.Intn(len(key))
		bytes = append(bytes, key[index])
	}
	return bytes
}

// 两个随机字节码
func twoBytes() []byte {
	// 获取两个随机的 ascii 码值
	key := GetOption().Key
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, 0)
	for i := 0; i < 2; i++ {
		index := r.Intn(len(key))
		bytes = append(bytes, key[index])
	}
	return bytes
}

// 乱序给定的字节数组
func ShuffleBytes(slice []byte) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
