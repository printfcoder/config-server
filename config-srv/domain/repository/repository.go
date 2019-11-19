package repository

import (
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/micro-in-cn/config-server/config-srv/config"
	"github.com/micro/go-micro/util/log"
)

var (
	once    sync.Once
	mysqlDB *gorm.DB
)

// Init 初始化数据库
func Init() {
	log.Infof("[Init] init db")

	once.Do(func() {
		if config.GetDBConfig().GetDialect() == "mysql" {
			initMysql()
		}
	})

	log.Infof("[Init] init db done")
}

// GetDB 获取db
func GetDB() *gorm.DB {
	return mysqlDB
}
