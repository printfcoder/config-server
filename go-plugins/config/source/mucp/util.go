package mucp

import (
	"github.com/micro/go-micro/config/source"
	"time"

	proto "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
)

func toChangeSet(c *proto.ChangeSet) *source.ChangeSet {
	return &source.ChangeSet{
		Data:      c.Data,
		Checksum:  c.Checksum,
		Format:    c.Format,
		Timestamp: time.Unix(c.Timestamp, 0),
		Source:    c.Source,
	}
}
