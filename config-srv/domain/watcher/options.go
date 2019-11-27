package watcher

import "github.com/micro-in-cn/config-server/config-srv/domain/dto"

type Options struct {
	UpdateChan <-chan *NSUpdate
}

type Option func(o *Options)

type NSUpdate struct {
	*dto.NSUpdate
	NewestDataBytes []byte
}

func WithUpdateChan(update <-chan *NSUpdate) Option {
	return func(o *Options) {
		o.UpdateChan = update
	}
}
