package repository

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/micro-in-cn/config-server/config-srv/config"
	"github.com/micro/go-micro/util/log"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

// Init 初始化数据库
func Init() {
	log.Infof("[Init] init db")

	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db inted")
		log.Logf(err.Error())
		return
	}

	if config.GetDBConfig().GetDialect() == "mysql" {
		initMysql()
	}

	log.Infof("[Init] init db done")
	inited = true
}

// GetDB 获取db
func GetDB() *sql.DB {
	return mysqlDB
}
