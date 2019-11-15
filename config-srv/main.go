package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-in-cn/config-server/config-srv/config"
	"github.com/micro-in-cn/config-server/config-srv/domain/repository"
	"github.com/micro-in-cn/config-server/config-srv/facade"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
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
