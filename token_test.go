package facote

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {

	//x, err := CreateToken(100003, 100004, int64(0))
	//if err != nil {
	//	println(fmt.Sprintf("%s", err))
	//	return
	//}
	//
	//// xYZA0L/bYC6CyDOZZF18vj026tFtTXXiQYHOd8EqfWs=
	//
	//println(fmt.Sprintf("Encrypt: %s", x))
	uc, qc, err := CheckToken("CQmzs8xf/hSSugETy7d5QUM2WEmI7k9TmUg7tW1EUFbVW6orJPiYPv7+XEIn/zU=")
	if err != nil {
		println(fmt.Sprintf("error: %s", err))
		return
	}
	println(fmt.Sprintf("Successful: uc: %d, qc: %d", uc, qc))

}

func TestCheckToken(t *testing.T) {
	uc, qc, err := CheckToken("zFBOTz7tZQw3BEBLAm+7TeJoqvaELIYhnXs+Ne6+WqEqOrYAUN7hEQkkB2CYQq46")
	if err != nil {
		println(fmt.Sprintf("%s", err))
		return
	}
	println(fmt.Sprintf("Successful: uc: %d, qc: %d", uc, qc))
}

func TestKeyBytes(t *testing.T) {

	bytes := make([]byte, 0)
	for i := 0; i < 256; i++ {
		bytes = append(bytes, byte(i))
	}
	println(fmt.Sprintf("%d", bytes))
	Shuffle(bytes)
	println(fmt.Sprintf("%d", bytes))

}

func Shuffle(slice []byte) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
