package db

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB
var dbUseList []string
var (
	TYPE_MYSQL = "mysql"
)

func InitDB(dbUse ...string) (err error) {
	if len(dbUse) <= 0 {
		return
	}
	dbUseList = dbUse
	var tmpErr error
	for _, t := range dbUseList {
		switch t {
		case TYPE_MYSQL:
			tmpErr = InitMysql()
		}
		if tmpErr != nil {
			err = errors.Join(err, tmpErr)
		}
	}
	return
}

func CloseDB() {
	for _, t := range dbUseList {
		switch t {
		case TYPE_MYSQL:
			// sqlDB, err := DB.DB()
			// if err == nil {
			// 	sqlDB.Close()
			// }
			fmt.Println("gorm's mysql is a pool connection, so you don't need to close it")
		}
	}
}
