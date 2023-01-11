package repo

import (
	"github.com/rodert/hepburn/config"
	"github.com/rodert/hepburn/internal/db/mysql"
	"github.com/rodert/hepburn/model/dao"
)

// mysql 链接初始化
func InitRepo(configure *config.Configuration) bool {
	return mysql.InitMysql(configure.MySQL)
}

// mysql 表结构初始化
func InitMySQLTable() (bool, error) {
	db, err := mysql.GetMySQL()
	if err != nil {
		return false, err
	}
	dao.NewUserProvider(db)

	return true, nil
}
