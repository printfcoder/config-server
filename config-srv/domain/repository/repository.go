package repository

import (
	"sync"

	"github.com/micro-in-cn/config-server/config-srv/config"
	"github.com/micro-in-cn/config-server/config-srv/domain/model"
	rgorm "github.com/micro-in-cn/config-server/config-srv/domain/repository/gorm"
	"github.com/micro/go-micro/util/log"
)

var (
	once   sync.Once
	repo   Repository
	models = []interface{}{&model.App{}, &model.Cluster{}, &model.Env{}, &model.Instance{}, &model.Item{}, &model.Namespace{}}
)

// Init 初始化数据库
func Init() {
	log.Infof("[Init] init db")

	once.Do(func() {
		switch config.GetDBConfig().GetORM() {
		case "gorm":
		default:
			repo = rgorm.NewGROMRepo(
				rgorm.WithDialect(config.GetDBConfig().GetDialect()),
				rgorm.WithAutoMigrate(config.GetDBConfig().GetGORM().GetAutoMigrate()),
				rgorm.WithLogMode(config.GetDBConfig().GetGORM().GetLogMode()),
				rgorm.WithSingularTable(true),
				// todo or pg?
				rgorm.WithURL(config.GetDBConfig().GetMysql().GetURL()),
				rgorm.WithMaxIdleConns(config.GetDBConfig().GetMysql().GetMaxIdleConnection()),
				rgorm.WithMaxOpenConns(config.GetDBConfig().GetMysql().GetMaxOpenConnection()),
				rgorm.WithAutoMigrateModels(models),
			)
		}
	})

	log.Infof("[Init] init db done")
}

func Repo() Repository {
	return repo
}
