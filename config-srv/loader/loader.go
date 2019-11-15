package loader

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"

	proto "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
)

//nolint
func loadConfigFile() (err error) {
	// watch changes
	watcher, err := config.Watch()
	if err != nil {
		log.Fatalf("[loadConfigFile] start watching files error，%s", err)
		return err
	}

	go func() {
		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatalf("[loadConfigFile] watch files error，%s", err)
				return
			}

			log.Logf("[loadConfigFile] file change， %s", string(v.Bytes()))
		}
	}()

	return
}

func GetConfig(appName string) *proto.ChangeSet {
	bytes := config.Get(appName).Bytes()

	log.Logf("[getConfig] appName，%s", string(bytes))
	return &proto.ChangeSet{
		Data:      bytes,
		Checksum:  fmt.Sprintf("%x", md5.Sum(bytes)),
		Format:    "yml",
		Source:    "file",
		Timestamp: time.Now().Unix()}
}

func ParsePath(path string) (appName string) {
	paths := strings.Split(path, "/")

	if paths[0] == "" && len(paths) > 1 {
		return paths[1]
	}

	return paths[0]
}
