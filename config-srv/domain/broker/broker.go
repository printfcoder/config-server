package broker

import (
	"encoding/json"
	"sync"

	"github.com/micro-in-cn/config-server/config-srv/domain/dto"
	"github.com/micro/go-micro/broker"
	"github.com/prometheus/common/log"
)

var (
	once sync.Once
)

type Broker interface {
	Run()
	Pub(update *dto.NSUpdate) error
}

type bk struct {
	opts Options
	// mu   sync.Mutex
	b broker.Broker
}

func (b *bk) Run() {
	once.Do(func() {
		if err := b.sub(); err != nil {
			panic(err)
		}
	})
}

func (b *bk) sub() error {
	_, err := b.b.Subscribe(b.opts.UpdateTopic, func(p broker.Event) error {
		log.Debug("[sub] Received Body: %s, Header: %s\n", string(p.Message().Body), p.Message().Header)
		update := dto.NSUpdate{}
		err := json.Unmarshal(p.Message().Body, &update)
		if err != nil {
			log.Error("[sub] broker sub err: %s", err)
			return err
		}

		b.opts.UpdateEvt <- update

		return nil
	})

	log.Error("[sub] broker sub err: %s", err)
	return err
}

func (b *bk) Pub(update *dto.NSUpdate) error {
	// todo more supports of Marshal
	bytes, _ := json.Marshal(update)

	msg := &broker.Message{
		Header: map[string]string{
			// todo placeholder
			"todo": "placeholder",
		},
		Body: bytes,
	}
	return b.b.Publish(b.opts.UpdateTopic, msg)
}

func NewBroker(opts ...Option) Broker {
	var options Options
	for _, o := range opts {
		o(&options)
	}

	return &bk{
		opts: options,
	}
}
