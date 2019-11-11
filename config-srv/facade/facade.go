package facade

import (
	protoS "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
	protoE "github.com/micro-in-cn/config-server/proto/entry"
	"github.com/micro/go-micro/server"
)

func RegisterHandlers(server server.Server) error {
	_ = protoS.RegisterSourceHandler(server, new(source))
	_ = protoE.RegisterEntryHandler(server, new(entry))
	return nil
}
