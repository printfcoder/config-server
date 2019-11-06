package db

import (
	"database/sql"

	"github.com/micro-in-cn/config-server/srv/config"
	"github.com/micro/go-micro/util/log"
)

func initMysql() {
	log.Infof("[initMysql] init db: mysql")

	var err error

	mysqlDB, err = sql.Open("mysql", config.GetDBConfig().GetMysql().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	mysqlDB.SetMaxOpenConns(config.GetDBConfig().GetMysql().GetMaxOpenConnection())
	mysqlDB.SetMaxIdleConns(config.GetDBConfig().GetMysql().GetMaxIdleConnection())

	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Infof("[initMysql] init db: mysql has inited done")
}
