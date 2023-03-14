package util

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	Root = "root"
)

var (
	RootConfig ConfigRoot
)

type ConfigRoot struct {
	SpeechKey    string `mapstructure:"SpeechKey"`
	SpeechRegion string `mapstructure:"SpeechRegion"`
	RpcPort      int16  `mapstructure:"RpcPort"`
	HttpPort     int16  `mapstructure:"HttpPort"`
}

func InitConfig(configFile string) error {
	viper.SetConfigFile(configFile)
	err := ReadConfig()
	if err != nil {
		return fmt.Errorf("reac config error, %s", err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		err := ReadConfig()
		if err != nil {
			Logger.Errorf("hot load config err, %s", err.Error())
			return
		}
	})
	viper.WatchConfig()
	return nil
}

func ReadConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config file error, %s", err.Error())
	}

	err := viper.Unmarshal(&RootConfig)
	if err != nil {
		return fmt.Errorf("read config error, %s", err.Error())
	}

	Logger.Infof("read config: %d, %d, %s, %s",
		RootConfig.RpcPort, RootConfig.HttpPort,
		RootConfig.SpeechRegion, RootConfig.SpeechKey)
	return nil
}
