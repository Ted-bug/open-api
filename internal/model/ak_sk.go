package model

import (
	"crypto"
	"crypto/hmac"
	"encoding/hex"

	"errors"

	"github.com/Ted-bug/open-api/internal/tool/mysql"
)

const T_AK_SK = "op_ak_sk"

type AkSk struct {
	Id         int32  `json:"id"`
	Uid        string `json:"uid"`
	Ak         string `json:"ak"`
	Sk         string `json:"sk"`
	Show       string `json:"show"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

// IsAkSkExist 检查指定的Access Key (ak) 是否存在
// 参数:
//
//	ak string - 需要检查的Access Key
//
// 返回值:
//
//	bool - 如果指定的Access Key存在，则返回true；否则返回false
func IsAkSkExist(ak string) bool {
	var aksk AkSk
	// 通过MySQL数据库查询指定Access Key是否存在
	if err := mysql.DB.Table(T_AK_SK).Where("ak=?", ak).Find(&aksk).Error; err != nil {
		return false // 如果查询出错，认为Access Key不存在
	}
	return true // 无错误发生，表示Access Key存在
}

func GetSk(ak string) (string, error) {
	var aksk AkSk
	if err := mysql.DB.Table(T_AK_SK).Where("ak=?", ak).Find(&aksk).Error; err != nil {
		return "", err
	}
	return aksk.Sk, nil
}

// CreateSign 使用密钥和时间戳生成签名
// 参数:
//
//	sk string - 私钥，用于签名计算的密钥
//	timestamp string - 时间戳，用于签名的内容
//
// 返回值:
//
//	string - 生成的签名字符串
//	error - 如果在签名过程中遇到错误，则返回错误信息
func CreateSign(sk string, timestamp string) (string, error) {
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

// CheckSign 用于验证签名是否有效
// 参数：
//
//	sk: 私钥，用于生成签名的密钥
//	timestamp: 时间戳，签名中包含的时间信息
//	sign: 待验证的签名字符串
//
// 返回值:
// bool - 验证结果，true 表示签名有效，false 表示签名无效。
// error - 验证过程中遇到的错误，如果验证通过则为 nil。
func CheckSign(sk, timestamp, sign string) (bool, error) {
	var deSign, cSign []byte
	var tmpSign string
	var err error

	// 将签名字符串从十六进制解码
	if deSign, err = hex.DecodeString(sign); err != nil {
		return false, err
	}

	// 使用私钥和时间戳尝试创建一个签名，以进行比对
	if tmpSign, err = CreateSign(sk, timestamp); err != nil {
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
