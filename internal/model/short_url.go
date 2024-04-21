package model

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/Ted-bug/open-api/internal/tool/mysql"
)

const T_SHORT_URL = "op_short_url"

type ShortUrl struct {
	Id         int32  `json:"id"`
	Url        string `json:"url"`
	Short      string `json:"short"`
	Hash       string `json:"hash"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

func IsUrlExist(Url string) (string, bool) {
	hasher := md5.New()
	if _, err := hasher.Write([]byte(Url)); err != nil {
		return "", false
	}
	urlMd5 := hex.EncodeToString(hasher.Sum(nil))
	var short ShortUrl
	if err := mysql.DB.Table(T_SHORT_URL).Where("hash=?", urlMd5).Find(&short).Error; err == nil {
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
	var short ShortUrl
	short.Url = Url
	short.Hash = hex.EncodeToString(hasher.Sum(nil))
	short.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	short.Short = generateShortNumber()
	short.Status = 1
	if err := mysql.DB.Table(T_SHORT_URL).Create(&short).Error; err != nil {
		return "", err
	}
	return short.Short, nil
}

// 生成短链接号
func generateShortNumber() string {
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
