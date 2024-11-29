package user_service

import (
	"crypto"
	"crypto/hmac"
	"encoding/hex"
	"errors"

	"github.com/Ted-bug/open-api/internal/model"
	"github.com/Ted-bug/open-api/internal/pkg/db"
	"github.com/jinzhu/gorm"
)

// GetUserSkByAk 检查指定的Access Key (ak) 是否存在
// 参数:
//
//	ak string - 需要检查的Access Key
//
// 返回值:
//
//	string - 有记录，返回sk
//
// error - 如果在查询过程中遇到错误，则返回错误信息
func GetUserSkByAk(ak string) (string, error) {
	// 通过MySQL数据库查询指定Access Key是否存在
	var aksk model.AkSk
	if err := db.DB.Where("ak=?", ak).Where("status=?", 1).First(&aksk).Error; err != nil || gorm.IsRecordNotFoundError(err) {
		return "", errors.New("查询出错")
	}
	return aksk.Sk, nil // 无错误发生，表示Access Key存在
}

// CreateSignByAkSk 使用密钥和时间戳生成签名
// 参数:
//
//	sk string - 私钥，用于签名计算的密钥
//	timestamp string - 时间戳，用于签名的内容
//
// 返回值:
//
//	string - 生成的签名字符串
//	error - 如果在签名过程中遇到错误，则返回错误信息
func CreateSignByAkSk(sk string, timestamp string) (string, error) {
	// 检查SHA256哈希函数是否可用
	if !crypto.SHA256.Available() {
		return "", errors.New("hash function is not available")
	}
	// 初始化HMAC-SHA1加密器，设置加密密钥
	hasher := hmac.New(crypto.SHA1.New, []byte(sk))
	// 向加密器中写入时间戳数据
	if _, err := hasher.Write([]byte(timestamp)); err != nil {
		return "", err
	}
	// 计算签名并返回
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// CheckSignByAkSk 用于验证签名是否有效
// 参数：
//
//	sk: 私钥，用于生成签名的密钥
//	timestamp: 时间戳，签名中包含的时间信息
//	sign: 待验证的签名字符串
//
// 返回值:
// bool - 验证结果，true 表示签名有效，false 表示签名无效。
// error - 验证过程中遇到的错误，如果验证通过则为 nil。
func CheckSignByAkSk(sk, timestamp, sign string) (bool, error) {
	var deSign, cSign []byte
	var tmpSign string
	var err error

	// 将签名字符串从十六进制解码
	if deSign, err = hex.DecodeString(sign); err != nil {
		return false, err
	}
	// 使用私钥和时间戳尝试创建一个签名，以进行比对
	if tmpSign, err = CreateSignByAkSk(sk, timestamp); err != nil {
		return false, err
	}
	// 将新创建的签名字符串从十六进制解码
	if cSign, err = hex.DecodeString(tmpSign); err != nil {
		return false, err
	}
	// 比对解码后的签名是否一致，不一致则验证失败
	if !hmac.Equal(deSign, cSign) {
		return false, errors.New("the sign is not avaliable")
	}
	// 签名验证通过，返回 true
	return true, nil
}
