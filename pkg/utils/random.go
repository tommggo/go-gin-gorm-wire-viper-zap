package utils

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var (
	defaultRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	orderSeq    uint32 // 订单号序列计数器
)

// RandomString 生成指定长度的随机字符串
// chars 参数指定字符集，如果为空则使用默认字符集（数字+字母）
func RandomString(length int, chars ...string) string {
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if len(chars) > 0 {
		charset = chars[0]
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[defaultRand.Intn(len(charset))]
	}
	return string(b)
}

// RandomInt 生成指定范围内的随机整数 [min, max)
func RandomInt(min, max int) int {
	return min + defaultRand.Intn(max-min)
}

// RandomDigits 生成指定长度的随机数字字符串
func RandomDigits(length int) string {
	return RandomString(length, "0123456789")
}

// GenerateOrderID 生成订单号
// 格式：年月日时分秒(14位) + 随机序号(4位)
// 示例：20250111143022-0001
func GenerateOrderID() string {
	// 获取当前时间
	now := time.Now()

	// 原子递增，确保线程安全
	seq := atomic.AddUint32(&orderSeq, 1)
	// 取余确保不超过4位
	seq = seq % 10000

	// 格式化订单号
	return fmt.Sprintf("%s%04d",
		now.Format("20060102150405"),
		seq,
	)
}
