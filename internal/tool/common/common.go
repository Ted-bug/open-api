package common

import (
	"math/rand"
	"time"
)

// 随机生成小时
func RandHour(l int64) time.Duration {
	return time.Duration(l)*time.Hour + time.Duration(rand.Int63n(10))*time.Minute
}

// 随机生成分钟
func RandMinute(l int64) time.Duration {
	return time.Duration(l)*time.Minute + time.Duration(rand.Int63n(10))*time.Second
}

// 随机生成秒
func RandSecond(l int64) time.Duration {
	return time.Duration(l)*time.Second + time.Duration(rand.Int63n(10))*time.Millisecond
}
