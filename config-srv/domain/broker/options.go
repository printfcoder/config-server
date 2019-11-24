package broker

import (
	"github.com/micro-in-cn/config-server/config-srv/domain/dto"
	"github.com/micro/go-micro/broker"
)

type Options struct {
	Broker      broker.Broker
	UpdateEvt   chan dto.NSUpdate
	UpdateTopic string
}

type Option func(o *Options)

func WithUpdateEvt(b chan dto.NSUpdate) Option {
	return func(o *Options) {
		o.UpdateEvt = b
	}
}

func WithBroker(b broker.Broker) Option {
	return func(o *Options) {
		o.Broker = b
	}
}

func WithUpdateTopic(updateTopic string) Option {
	return func(o *Options) {
		o.UpdateTopic = updateTopic
	}
}
