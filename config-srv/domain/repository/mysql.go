package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/micro-in-cn/config-server/config-srv/config"
	"github.com/micro-in-cn/config-server/config-srv/domain/model"
	"github.com/micro/go-micro/util/log"
)

func initMysql() {
	log.Infof("[initMysql] init db: mysql")

	var err error
	if mysqlDB, err = gorm.Open("mysql", config.GetDBConfig().GetMysql().GetURL()); err != nil {
		log.Fatal(err)
	}
	mysqlDB.SingularTable(true)
	mysqlDB.DB().SetMaxOpenConns(100)
	mysqlDB.DB().SetMaxIdleConns(10)
	mysqlDB.LogMode(true)

	if err = mysqlDB.AutoMigrate(&model.App{}, &model.Cluster{}, &model.Env{}, &model.Instance{}, &model.Item{}, &model.Namespace{}).Error; err != nil {
		_ = mysqlDB.Close()
		log.Fatal(err)
	}

	if err = mysqlDB.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	log.Infof("[initMysql] init db: mysql has inited done")
}
