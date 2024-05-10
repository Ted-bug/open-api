package common

import (
	"math/rand"
	"net/http"
	"regexp"
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

// 检查url是否可达：Get请求方式
func IsUrlActive(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

// 检查url格式是否正确
func IsValidURL(url string) bool {
	// 正则表达式来匹配大部分标准URL
	// 这个正则表达式可能不能完全匹配所有有效的URL，但足够用于大多数情况
	regex := `^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:\/?#[\]@!\$&'\(\)\*\+,;=.]+$`
	matched, err := regexp.MatchString(regex, url)
	if err != nil {
		// 如果出现错误，则认为URL无效
		return false
	}
	return matched
}
