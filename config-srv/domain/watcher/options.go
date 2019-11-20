package watcher

type Options struct {
	UpdateChan chan *NSUpdate
}

type Option func(o *Options)

type NSUpdate struct {
	AppId       string
	Cluster     string
	Namespace   string
	Env         string
	NewestBytes []byte
}

func WithUpdateChan(update chan *NSUpdate) Option {
	return func(o *Options) {
		o.UpdateChan = update
	}
}
