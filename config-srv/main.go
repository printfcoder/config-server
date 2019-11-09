package main

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-in-cn/config-server/config-srv/config"
	"github.com/micro-in-cn/config-server/config-srv/facade"
	"github.com/micro-in-cn/config-server/config-srv/repository"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

var (
	mux sync.RWMutex
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.config"),
		micro.WrapCall(),
	)

	_ = facade.RegisterHandlers(service.Server())

	service.Init(micro.Action(func(ctx *cli.Context) {
		config.Init()
		repository.Init()
	}))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
