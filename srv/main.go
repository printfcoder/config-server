package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/micro/go-micro"
	"strings"
	"sync"
	"time"

	proto "github.com/micro-in-cn/config-server/go-plugins/config/source/mucp/proto"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
)

var (
	mux        sync.RWMutex
	configMaps = make(map[string]*proto.ChangeSet)
	apps       = []string{"micro", "extra"}
)

type Source struct{}

func main() {
	// load config files
	err := loadConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	// 新建服务
	service := micro.NewService(
		micro.Name("go.micro.config"),
	)

	// 注册服务
	proto.RegisterSourceHandler(service.Server(), new(Source))

	service.Init()

	if err := service.Run(); err != nil {
		panic(err)
	}
}

func (s Source) Read(ctx context.Context, req *proto.ReadRequest, rsp *proto.ReadResponse) (err error) {
	appName := parsePath(req.Path)
	switch appName {
	case "micro", "extra":
		rsp.ChangeSet = getConfig(appName)
		return
	default:
		err = fmt.Errorf("[Read] the first path is invalid")
		return
	}

	return
}

func (s Source) Watch(ctx context.Context, req *proto.WatchRequest, server proto.Source_WatchStream) (err error) {
	appName := parsePath(req.Path)
	rsp := &proto.WatchResponse{
		ChangeSet: getConfig(appName),
	}
	if err = server.SendMsg(rsp); err != nil {
		log.Logf("[Watch] watch files error，%s", err)
		return err
	}

	return
}

func loadConfigFile() (err error) {
	for _, app := range apps {
		if err := config.Load(file.NewSource(
			file.WithPath("D:\\GOPATH\\src\\github.com\\micro-in-cn\\config-server\\srv\\conf\\" + app + ".yml"),
		)); err != nil {
			log.Fatalf("[loadConfigFile] load files error，%s", err)
			return err
		}
	}

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

func getConfig(appName string) *proto.ChangeSet {
	bytes := config.Get(appName).Bytes()

	log.Logf("[getConfig] appName，%s", string(bytes))
	return &proto.ChangeSet{
		Data:      bytes,
		Checksum:  fmt.Sprintf("%x", md5.Sum(bytes)),
		Format:    "yml",
		Source:    "file",
		Timestamp: time.Now().Unix()}
}

func parsePath(path string) (appName string) {
	paths := strings.Split(path, "/")

	if paths[0] == "" && len(paths) > 1 {
		return paths[1]
	}

	return paths[0]
}
