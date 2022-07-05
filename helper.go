package facote

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type CryptoToken struct {
	Key string `json:"key"`
	Ts  int64  `json:"ts"`
	Ps  string `json:"ps"`
}

type CryptoKey struct {
	Token string
	Key   []byte
	Unkey []byte
}

// 将字节码转换为 base64
func (ck *CryptoKey) Encrypt(data []byte) []byte {
	// 乱序
	for i := 0; i < len(data); i++ {
		index := data[i]
		data[i] = ck.Key[index]
	}
	return data
}

func (ck *CryptoKey) Decrypt(data []byte) []byte {
	// 正序
	for i := 0; i < len(data); i++ {
		index := data[i]
		data[i] = ck.Unkey[index]
	}
	return data
}

type CryptoHelper struct {
	count int // 插入随机数的数量
	key   []byte
	unkey []byte
	ps    string
	td    int64
}

func CreateHelper(count int, ps string, key []byte) *CryptoHelper {
	// 计算反向 Key
	bytes := make([]byte, 128)
	for i := 0; i < len(key); i++ {
		index := key[i]
		bytes[index] = byte(i)
	}
	// 返回 Helper
	return &CryptoHelper{
		count: count,
		key:   key,
		unkey: bytes,
		ps:    ps,
	}
}

// 时差可后设置
func (ch *CryptoHelper) SetTd(td int64) {
	ch.td = td
}

func (ch *CryptoHelper) CreateToken() (*CryptoKey, error) {
	// key
	key := make([]byte, 0)
	for i := 0; i < 128; i++ {
		key = append(key, byte(i))
	}
	ShuffleBytes(key)
	// unkey
	unkey := make([]byte, 128)
	for i := 0; i < len(key); i++ {
		index := key[i]
		unkey[index] = byte(i)
	}
	ck := &CryptoKey{
		Key:   key,
		Unkey: unkey,
	}
	// 创建 token
	keybase := base64.StdEncoding.EncodeToString(key)
	ct := &CryptoToken{}
	ct.Key = keybase
	ct.Ps = ch.ps
	timestamp := time.Now().Unix()
	ct.Ts = timestamp + ch.td
	// 转换为 Json 字符串
	data, err := json.Marshal(ct)
	if err != nil {
		return nil, err
	}
	encryptData := ch.Encrypt(data)
	token := base64.StdEncoding.EncodeToString(encryptData)
	ck.Token = token
	return ck, nil
}

// Token 检测，可选择是否校验时间差
func (ch *CryptoHelper) CheckToken(token string, checked bool) (*CryptoKey, error) {
	// Base64 解密
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Base64 decode error: %s", err))
	}
	// 解密
	decryptData := ch.Decrypt(data)
	// Json 转换为对象
	var ct *CryptoToken
	err = json.Unmarshal(decryptData, &ct)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Json unmarshal error: %s", err))
	}
	// 校验密码
	if ct.Ps != ch.ps {
		return nil, errors.New("check token ps error")
	}
	// 校验时差
	if checked {
		when := ct.Ts
		now := time.Now().Unix()
		if now-when > 30 || now-when < -10 {
			return nil, errors.New("check token timestamp error")
		}
	}
	// 返回
	key, err := base64.StdEncoding.DecodeString(ct.Key)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Base64 decode error: %s", err))
	}
	ck := &CryptoKey{
		Token: token,
		Key:   key,
	}
	return ck, nil
}

// 对字节码数组加密
func (ch *CryptoHelper) Encrypt(data []byte) []byte {
	bytes := make([]byte, 0)
	// 添加一个随机长度的头部，再每个正确字符前添加两个随机字符
	head := ch.header()
	bytes = append(bytes, head...)
	for i := 0; i < len(data); i++ {
		bytes = append(bytes, ch.randBytes()...)
		bytes = append(bytes, data[i])
	}
	// 进行乱序
	for i := 0; i < len(bytes); i++ {
		index := bytes[i]
		bytes[i] = ch.key[index]
	}
	return bytes
}

// 对字节码数组解密
func (ch *CryptoHelper) Decrypt(data []byte) []byte {
	for i := 0; i < len(data); i++ {
		index := data[i]
		data[i] = ch.unkey[index]
	}
	bytes := make([]byte, 0)
	var count int
	for i, v := range data {
		if i == 3 {
			count = int(v)
		}
		if (i-count-3) > 0 && (i-count-3)%(ch.count+1) == 0 {
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
func (ch *CryptoHelper) header() []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, 0)
	for i := 0; i < 3; i++ {
		index := r.Intn(len(ch.key))
		bytes = append(bytes, ch.key[index])
	}
	count := r.Intn(30)
	bytes = append(bytes, byte(count))
	for i := 0; i < count; i++ {
		index := r.Intn(len(ch.key))
		bytes = append(bytes, ch.key[index])
	}
	return bytes
}

// 随机字节码
func (ch *CryptoHelper) randBytes() []byte {
	// 获取 count 个随机的 ascii 码值
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, 0)
	for i := 0; i < ch.count; i++ {
		index := r.Intn(len(ch.key))
		bytes = append(bytes, ch.key[index])
	}
	return bytes
}
