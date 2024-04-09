package mysql

import (
	"github.com/Ted-bug/open-api/internal/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitMysql() error {
	var err error
	DB, err = NewMysql("mysql", config.AppConfig.Mysql)
	if err != nil {
		return err
	}
	err = DB.DB().Ping()
	if err != nil {
		return err
	}
	return nil
}

// 获取新的数据库连接
func NewMysql(section string, option config.Mysql) (*gorm.DB, error) {
	dns := GetMySqlDSN(section, option)
	newDb, err := gorm.Open(section, dns)
	if err != nil {
		return nil, err
	}
	return newDb, nil
}

// 组装数据库连接链接
func GetMySqlDSN(section string, option config.Mysql) string {
	dsn := option.User + ":" + option.Password
	dsn += "@(" + option.Host + ":" + option.Port + ")/"
	dsn += option.DbName + "?charset=" + option.Charset
	dsn += "&parseTime=True&loc=Local"
	return dsn
}

func CloseMysql() {
	DB.Close()
}
