package shorturlservice

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/Ted-bug/open-api/internal/model"
	"github.com/Ted-bug/open-api/internal/tool/mysql"
	"github.com/Ted-bug/open-api/internal/tool/redis"
)

const (
	// 短链接长度
	SHORT_KEY = "short-list:"
)

// 判断短链接是否存在
// Url 长链接
// 返回短链接号和是否存在
func IsUrlExist(Url string) (string, bool) {
	hasher := md5.New()
	if _, err := hasher.Write([]byte(Url)); err != nil {
		return "", false
	}
	urlMd5 := hex.EncodeToString(hasher.Sum(nil))
	key := SHORT_KEY + urlMd5
	if has, err := redis.RedisClient.Exists(key).Result(); err == nil && has != 0 {
		if v, err := redis.RedisClient.Get(key).Result(); err == nil {
			return v, true
		}
	}
	var short model.ShortUrl
	if err := mysql.DB.Where("hash=?", urlMd5).Find(&short).Error; err == nil {
		redis.RedisClient.Set(key, short.Short, 10*time.Second).Result()
		return short.Short, true
	}
	return "", false
}

// 创建短链接
func CreateShortUrl(Url string) (string, error) {
	hasher := md5.New()
	if _, err := hasher.Write([]byte(Url)); err != nil {
		return "", err
	}
	var short model.ShortUrl
	short.Url = Url
	short.Hash = hex.EncodeToString(hasher.Sum(nil))
	short.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	short.Short = generateNumber()
	short.Status = 1
	if err := mysql.DB.Create(&short).Error; err != nil {
		return "", err
	}
	key := SHORT_KEY + short.Hash
	redis.RedisClient.Set(key, short.Short, 10*time.Second).Result()
	return short.Short, nil
}

// 生成短链接号
func generateNumber() string {
	base62Chars := []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	number := rand.Int63()

	short, remainder := []byte{}, int64(0)
	for number > 0 {
		number, remainder = number/64, number%64
		short = append([]byte{base62Chars[remainder]}, short...)
	}
	return string(short)
}
