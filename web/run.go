package web

import (
	"github.com/rodert/hepburn/config"
	"github.com/rodert/hepburn/internal/db/repo"
	"github.com/rodert/hepburn/internal/server"
	"github.com/sirupsen/logrus"
)

func Run() {

	defer func() {
		if err := recover(); err != nil {
			logrus.Error(err)
		}
	}()

	// TODO 初始化连接
	// if err := log.NewLog(config.Configure.Log); err != nil {
	// 	return
	// }

	if ok := repo.InitRepo(&config.Configure); !ok {
		return
	}
	if ok, err := repo.InitMySQLTable(); !ok {
		logrus.Error(err)
		return
	}

	server.Serv(Router())
}
