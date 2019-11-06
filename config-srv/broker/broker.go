package broker

import "github.com/micro/go-micro/broker"

var (
	DefaultConfigChangeTopic = "go.micro.broker.config.change"
	bk                       broker.Broker
)

func Init(bk broker.Broker) {
	bk = bk
}

func subscribe() {

}

func publish() {

}
