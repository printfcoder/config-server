package facade

import (
	"github.com/micro-in-cn/config-server/config-srv/domain/service"
	protoS "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
	protoE "github.com/micro-in-cn/config-server/proto/entry"
	"github.com/micro/go-micro/server"
)

var (
	se service.Service
)

func Init() {
	se = service.GetService()
}

func RegisterHandlers(server server.Server) error {
	_ = protoS.RegisterSourceHandler(server, new(source))
	_ = protoE.RegisterEntryHandler(server, new(entry))
	return nil
}
