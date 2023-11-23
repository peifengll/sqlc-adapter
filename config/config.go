package config

import (
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"log"
	"sync"
)

var (
	c    *TomConfig
	once sync.Once
)

type TomConfig struct {
	AppName string
	Mysql   MySQLConfig
	Cache   cache.CacheConf
}

type MySQLConfig struct {
	Host        string
	DbName      string
	Password    string
	Port        int
	TablePrefix string
	User        string
}

func InitConfig() error {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	viper.AddConfigPath("D:\\code\\go\\Projects\\sqlc-adapter\\config")
	err := viper.ReadInConfig()
	if nil != err {
		return err
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}
	return nil
}
func GetConfig() *TomConfig {
	once.Do(func() {
		err := InitConfig()
		if err != nil {
			log.Fatal("配置加载错误")
			return
		}

	})
	return c
}
