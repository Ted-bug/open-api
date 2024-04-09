package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	AppDebug bool `mapstructure:"app_debug"`
	Mysql    struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DbName   string `mapstructure:"db_name"`
		Charset  string `mapstructure:"charset"`
		Prefix   string `mapstructure:"prefix"`
	} `mapstructure:"mysql"`
	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Db       int    `mapstructure:"db"`
		Prefix   string `mapstructure:"prefix"`
	} `mapstructure:"redis"`
}

var AppConfig = &Config{}

func InitConfig() error {
	// 初始化配置
	configViper := viper.New()
	if home, err := os.UserHomeDir(); err == nil {
		configViper.AddConfigPath(home)
	}
	configViper.AddConfigPath("./config")
	configViper.SetConfigName("config")
	configViper.SetConfigType("yaml")
	if err := configViper.ReadInConfig(); err != nil {
		return err
	}
	if err := configViper.Unmarshal(AppConfig); err != nil {
		return err
	}
	configViper.OnConfigChange(func(in fsnotify.Event) {
		configViper.ReadInConfig()
		if err := configViper.Unmarshal(AppConfig); err != nil {
			fmt.Println("reload config failed, err:", err)
		} else {
			fmt.Println("reload config success")
		}
	})
	configViper.WatchConfig()
	return nil
}
