package domain

import (
	"github.com/micro-in-cn/config-server/config-srv/domain/broker"
	"github.com/micro-in-cn/config-server/config-srv/domain/dto"
	"sync"

	"github.com/micro-in-cn/config-server/config-srv/domain/repository"
	"github.com/micro-in-cn/config-server/config-srv/domain/service"
	"github.com/micro-in-cn/config-server/config-srv/domain/watcher"
	mBroker "github.com/micro/go-micro/broker"
)

var (
	d = &domain{}
	// Service pushes the updateData to update chan of the Watcher, then Watcher calls every Stream to send data
	// to the clients by SendMsg.
	updateNS       = make(chan *watcher.NSUpdate)
	updateNSForPub = make(chan *dto.NSUpdate)
	updatePubTopic = "go.micro.config_server.topic.update_ns"
)

type domain struct {
	once sync.Once
}

func (d *domain) run(mbk mBroker.Broker) {
	d.once.Do(func() {
		// todo make the Init method in an unified style.
		repository.Init()
		service.Init(repository.Repo(),
			updateNS,
			updateNSForPub,
		)

		// broker pubs and subs the update of Namespace for the sibling nodes,
		// so they can update cache immediately as soon as possible.
		bk := broker.NewBroker(
			broker.WithUpdateEvt(updateNSForPub),
			broker.WithUpdateTopic(updatePubTopic),
			broker.WithBroker(mbk),
		)
		bk.Run()

		// watcher listens the update of Namespace for the remote Watchers of the clients connecting this server.
		watcher.NewWatcher(watcher.WithUpdateChan(updateNS))
	})
}

func Init(bk mBroker.Broker) {
	d.run(bk)
}
