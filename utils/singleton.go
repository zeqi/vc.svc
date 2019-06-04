package utils

import (
	"sync"

	"github.com/micro/go-config"
	lutils "vc.libs/utils"
	"vc.svc/models"
)

func init() {
	GetInstance()
}

type singleton struct {
	Config models.Config
}

var instance *singleton
var once sync.Once

func InitConfig(configDirs ...string) models.Config {
	var baseDir = "config"
	if len(configDirs) > 0 {
		baseDir = configDirs[0]
	}

	configHander := lutils.ConfigFIle{BaseDir: baseDir}
	configHander.ConfigFileHandler()
	conf := models.Config{}
	config.LoadFile(configHander.Path)
	config.Scan(&conf)
	return conf
}

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
		instance.Config = InitConfig()
	})
	return instance
}

func GetConfig() models.Config {
	return instance.Config
}
