package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Mode   string `mapstructure:"mode"`
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
	Logger Logger `mapstructure:"logger"`
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
	Charset  string `mapstructure:"charset"`
	Prefix   string `mapstructure:"prefix"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	Prefix   string `mapstructure:"prefix"`
}

type Logger struct {
	Type       string `mapstructure:"type"`
	Path       string `mapstructure:"path"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
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
	configViper.WatchConfig()
	configViper.OnConfigChange(func(in fsnotify.Event) {
		if err := configViper.Unmarshal(AppConfig); err != nil {
			fmt.Println("reload config failed, err:", err)
		} else {
			fmt.Println("reload config success")
		}
	})
	fmt.Println("Config init success")
	return nil
}
