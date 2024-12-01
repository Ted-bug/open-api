package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/Ted-bug/open-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var levelMap = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"info":   logger.Info,
	"warn":   logger.Warn,
	"error":  logger.Error,
}

func InitMysql() error {
	var (
		err error
		dsn string
	)
	dsn = GetMySqlDSN(TYPE_MYSQL, config.AppConfig.Mysql)
	if dsn == "" {
		return errors.New("db's dns is empty")
	}
	DB, err = gorm.Open(mysql.New(
		mysql.Config{
			DSN:               dsn,
			DefaultStringSize: 256,
		}),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.LogLevel(levelMap[config.AppConfig.Gorm.LogLevel])),
		})
	if err != nil {
		return err
	}
	if sqlDB, err := DB.DB(); err == nil {
		sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(config.AppConfig.Mysql.ConnMaxIdleTime))
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.AppConfig.Mysql.ConnMaxLifetime))
		sqlDB.SetMaxIdleConns(config.AppConfig.Mysql.MaxIdleConns)
		sqlDB.SetMaxOpenConns(config.AppConfig.Mysql.MaxOpenConns)
	} else {
		fmt.Println("mysql init connect's config failed")
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
