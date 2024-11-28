package db

import (
	"errors"
	"fmt"

	"github.com/Ted-bug/open-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysql() error {
	var (
		err error
		dsn string
	)
	dsn = GetMySqlDSN(TYPE_MYSQL, config.AppConfig.Mysql)
	if dsn == "" {
		return errors.New("db's dns is empty")
	}
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	fmt.Println("mysql init success")
	return nil
}

// 组装数据库连接链接
func GetMySqlDSN(section string, option config.Mysql) string {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		option.User, option.Password,
		option.Host, option.Port,
		option.DbName, option.Charset,
	)
	return dsn
}
