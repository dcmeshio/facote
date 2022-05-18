package facote

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/dcmeshio/facote/common"
	"time"
)

type Token struct {
	Uc int   `json:"uc"`
	Qc int   `json:"qc"`
	Ts int64 `json:"ts"`
}

func CreateToken(uc, qc int, td int64) (string, error) {
	option := GetOption()
	// 创建结构体
	t := &Token{}
	t.Uc = uc
	t.Qc = qc
	timestamp := time.Now().Unix()
	t.Ts = timestamp + td
	// 加密为字符串
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	encryptData, err := common.Encrypt(data, []byte(option.Key))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptData), nil
}

func CheckToken(t string) (int, int, error) {
	data, err := base64.StdEncoding.DecodeString(t)
	if err != nil {
		return 0, 0, err
	}
	option := GetOption()
	decryptData, err := common.Decrypt(data, []byte(option.Key))
	if err != nil {
		return 0, 0, err
	}
	var dt *Token
	err = json.Unmarshal(decryptData, &dt)
	if err != nil {
		return 0, 0, err
	}
	when := dt.Ts
	now := time.Now().Unix()
	if now-when > 30 || now-when < -10 {
		return 0, 0, errors.New("check token timestamp error")
	}
	return dt.Uc, dt.Qc, nil
}
