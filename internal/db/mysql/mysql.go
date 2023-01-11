package mysql

import (
	"errors"
	"time"

	"github.com/rodert/hepburn/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetMySQL() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}
	return db, errors.New("db connect faild")
}

func InitMysql(config config.MySQL) bool {
	var err error
	for i := 0; i < 3; i++ {
		db, err = gorm.Open(mysql.Open(config.DNS), &gorm.Config{})
		if err == nil {
			return true
		}
		time.Sleep(time.Second * 5)
	}
	return false
}
