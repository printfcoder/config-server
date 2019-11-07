package facade

import (
	proto "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
	"github.com/micro/go-micro/server"
)

func RegisterHandlers(server server.Server) error {
	_ = proto.RegisterSourceHandler(server, new(source))
	return nil
}
