package common

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
)

func UUID() string {
	UUID := uuid.New()
	return strings.ReplaceAll(UUID.String(), "-", "")
}

// 12 位字符串随机
func Rstring12() string {
	uid := UUID()
	return uid[:10]

}

// 8 位数字随机
func Rint8() string {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(90000000)
	return fmt.Sprintf("%d", value+10000000)
}

// 获取数字以内随机数
func Rint(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
