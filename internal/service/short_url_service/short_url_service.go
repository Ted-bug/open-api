package shorturlservice

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/Ted-bug/open-api/internal/model"
	"github.com/Ted-bug/open-api/internal/tool/common"
	"github.com/Ted-bug/open-api/internal/tool/mysql"
	"github.com/Ted-bug/open-api/internal/tool/redis"
)

const (
	SHORT_KEY  = "short-list:"  // 短链缓存
	IGNORE_KEY = "ignore-list:" // 短转长，忽略不存在的短链
	LONG_KEY   = "long-list:"   // 长链缓存
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
	// 优先缓存读取
	key := SHORT_KEY + urlMd5
	if has, err := redis.RedisClient.Exists(key).Result(); err == nil && has != 0 {
		if v, err := redis.RedisClient.Get(key).Result(); err == nil {
			return v, true
		}
	}
	var short model.ShortUrl
	if err := mysql.DB.Where("hash=?", urlMd5).Find(&short).Error; err == nil {
		redis.RedisClient.Set(key, short.Surl, common.RandMinute(5)).Result()
		return short.Surl, true
	}
	return "", false
}

// 创建短链接
func ConvertLurl(Url string) (string, error) {
	hasher := md5.New()
	if _, err := hasher.Write([]byte(Url)); err != nil {
		return "", err
	}
	var short model.ShortUrl
	short.Lurl = Url
	short.Hash = hex.EncodeToString(hasher.Sum(nil))
	short.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	short.Surl = generateNumber()
	short.Status = 1
	if err := mysql.DB.Create(&short).Error; err != nil {
		return "", err
	}
	// 加入缓存
	key := SHORT_KEY + short.Hash
	redis.RedisClient.Set(key, short.Surl, common.RandMinute(5)).Result()
	// 将短链接移除忽略列表
	redis.RedisClient.Del(IGNORE_KEY + short.Surl).Result()
	return short.Surl, nil
}

// 生成短链接号
func generateNumber() string {
	base62Chars := []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	number := rand.Int63()

	short, remainder := []byte{}, int64(0)
	for number > 0 {
		number, remainder = number/64, number%64
		short = append([]byte{base62Chars[remainder]}, short...)
	}
	return string(short)
}

// 解析短链
func RevertSurl(s string) (string, bool) {
	// 判断是否在忽略列表
	if has, err := redis.RedisClient.Exists(IGNORE_KEY + s).Result(); err == nil && has != 0 {
		return "", false
	}
	if v, err := redis.RedisClient.Get(LONG_KEY + s).Result(); err == nil {
		return v, true
	}
	var short model.ShortUrl
	if err := mysql.DB.Where("surl=?", s).Find(&short).Error; err == nil {
		redis.RedisClient.Set(LONG_KEY+s, short.Lurl, common.RandMinute(15)).Result()
		return short.Lurl, true
	}
	// 添加忽略列表
	redis.RedisClient.Set(IGNORE_KEY+s, 1, common.RandHour(10)).Result()
	return "", false
}
