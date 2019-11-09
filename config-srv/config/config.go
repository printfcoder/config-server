package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
)

var (
	m      sync.RWMutex
	inited bool
	sp     = string(filepath.Separator)

	configMe = &allConfig{}
)

type allConfig struct {
	DB *dbConfig `json:"db"`
}

// Init 初始化配置
func Init() {
	log.Infof("[Init] init config")

	m.Lock()
	defer m.Unlock()

	if inited {
		log.Logf("[Init] config inited")
		return
	}

	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("."+sp, sp)))

	pt := filepath.Join(appPath, "resources")
	os.Chdir(appPath)

	if err := config.Load(file.NewSource(file.WithPath(pt + sp + "application.yml"))); err != nil {
		panic(err)
	}

	if err := config.Get("micro", "config_server").Scan(configMe); err != nil {
		panic(err)
	}

	inited = true

	log.Infof("[Init] config init has been done")
}

func GetDBConfig() (ret DBConfig) {
	return configMe.DB
}
