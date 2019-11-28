package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-in-cn/config-server/config-srv/config"
	"github.com/micro-in-cn/config-server/config-srv/domain"
	"github.com/micro-in-cn/config-server/config-srv/facade"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/server"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.config"),
		micro.WrapCall(),
	)

	configSrvStart(service.Server(), service.Options().Broker)
	service.Init()

	if err := service.Run(); err != nil {
		panic(err)
	}
}

func configSrvStart(server server.Server, bk broker.Broker) {
	_ = facade.RegisterHandlers(server)
	config.Init()
	domain.Init(bk)
}
