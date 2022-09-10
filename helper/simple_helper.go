package helper

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dcmeshio/facote"
)

type SimpleCryptoToken struct {
	Key string `json:"key"`
	Ps  string `json:"ps"`
}

type SimpleCryptoHelper struct {
	key   []byte
	unkey []byte
	ps    string // 设定一个固定长度的密码：
}

func CreateHelper(ps string, key []byte) *SimpleCryptoHelper {
	// 计算反向 Key
	bytes := make([]byte, 128)
	for i := 0; i < len(key); i++ {
		index := key[i]
		bytes[index] = byte(i)
	}
	// 返回 Helper
	return &SimpleCryptoHelper{
		key:   key,
		unkey: bytes,
		ps:    ps,
	}
}

func (sch *SimpleCryptoHelper) createKey() (*facote.CryptoKey, error) {
	// key
	key := make([]byte, 0)
	for i := 0; i < 128; i++ {
		key = append(key, byte(i))
	}
	facote.ShuffleBytes(key)
	// unkey
	unkey := make([]byte, 128)
	for i := 0; i < len(key); i++ {
		index := key[i]
		unkey[index] = byte(i)
	}
	ck := &facote.CryptoKey{
		Key:   key,
		Unkey: unkey,
	}
	// 创建 token
	keybase64 := base64.StdEncoding.EncodeToString(key) // 私钥 base64
	psData := []byte(sch.ps)
	psData = ck.Encrypt(psData)
	psbase64 := base64.StdEncoding.EncodeToString(psData) // 密码 私钥加密 + base64
	sct := &SimpleCryptoToken{
		Key: keybase64,
		Ps:  psbase64,
	}
	data, err := json.Marshal(sct)
	if err != nil {
		return nil, err
	}
	encryptData := sch.Encrypt(data) // token 公钥加密
	token := base64.StdEncoding.EncodeToString(encryptData)
	ck.Token = token
	return ck, nil
}

func (sch *SimpleCryptoHelper) CheckToken(token string) (*facote.CryptoKey, error) {
	// Base64 解密
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Base64 decode error: %s", err))
	}
	// 解密
	decryptData := sch.Decrypt(data)
	// Json 转换为对象
	var sct *SimpleCryptoToken
	err = json.Unmarshal(decryptData, &sct)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Json unmarshal error: %s", err))
	}
	// 获取私钥
	key, err := base64.StdEncoding.DecodeString(sct.Key)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Base64 decode error: %s", err))
	}
	unkey := make([]byte, 128)
	for i := 0; i < len(key); i++ {
		index := key[i]
		unkey[index] = byte(i)
	}
	ck := &facote.CryptoKey{
		Token: token,
		Key:   key,
		Unkey: unkey,
	}
	// 校验密码
	psBytes, err := base64.StdEncoding.DecodeString(sct.Ps)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Base64 decode error: %s", err))
	}
	psBytes = ck.Decrypt(psBytes)
	if sch.ps != string(psBytes) {
		return nil, errors.New("check token ps error")
	}
	return ck, nil
}

// 加密
func (sch *SimpleCryptoHelper) Encrypt(data []byte) []byte {
	// 乱序
	bytes := make([]byte, 0)
	for i := 0; i < len(data); i++ {
		index := data[i]
		bytes = append(bytes, sch.key[index])
	}
	return bytes
}

// 解密
func (sch *SimpleCryptoHelper) Decrypt(data []byte) []byte {
	// 正序
	bytes := make([]byte, 0)
	for i := 0; i < len(data); i++ {
		index := data[i]
		bytes = append(bytes, sch.unkey[index])
	}
	return bytes
}
