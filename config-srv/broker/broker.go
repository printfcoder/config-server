package broker

import "github.com/micro/go-micro/broker"

var (
	DefaultConfigChangeTopic = "go.micro.broker.config.change"
	bk                       broker.Broker
)

func Init(broker broker.Broker) {
	bk = broker
}

func subscribe() {

}

func publish() {

}
