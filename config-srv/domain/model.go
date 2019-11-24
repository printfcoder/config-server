package domain

import (
	"github.com/micro-in-cn/config-server/config-srv/source/mysql"
	"sync"

	"github.com/micro-in-cn/config-server/config-srv/domain/repository"
	"github.com/micro-in-cn/config-server/config-srv/domain/service"
	"github.com/micro-in-cn/config-server/config-srv/domain/watcher"
)

var (
	d      = &domain{}
	update = make(chan *watcher.NSUpdate)
)

type domain struct {
	once sync.Once
}

func (d *domain) run() {
	d.once.Do(func() {
		repository.Init()
		service.Init(repository.Repo(), update)
		watcher.NewWatcher()

		mysqlSource := mysql.NewSource(mysql)
	})
}

func Init() {
	d.run()
}
