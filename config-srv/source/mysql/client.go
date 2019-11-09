package mysql

import (
	"github.com/micro/go-micro/config/source"
)

type Client struct {
	Service
}

func (c *Client) Query(path string) (set *source.ChangeSet, err error) {
	return nil, nil
}
