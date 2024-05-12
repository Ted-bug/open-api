package config

import (
	"fmt"
	"os"

	"github.com/Ted-bug/open-api/internal/constants"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
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
	configViper.AddConfigPath(constants.CONFPATH)
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

// 生成配置文件
func CreateConfig() error {
	config := Config{
		Mode: "debug",
		Host: "0.0.0.0",
		Port: "8080",
		Mysql: Mysql{
			Host:     "127.0.0.1",
			Port:     "3306",
			User:     "root",
			Password: "root",
			DbName:   "open_api",
			Charset:  "utf8mb4",
			Prefix:   "open_",
		},
		Redis: Redis{
			Host:     "127.0.0.1",
			Port:     "6379",
			Password: "",
			Db:       0,
			Prefix:   "open_",
		},
		Logger: Logger{
			Type:       "file",
			Path:       "./logs/",
			Filename:   "open_api.log",
			MaxSize:    10,
			MaxAge:     30,
			MaxBackups: 7,
		},
	}
	var cMap map[string]any
	mapstructure.Decode(config, &cMap)

	configViper := viper.New()
	for k1, v1 := range cMap {
		if _, ok := v1.(map[string]any); ok {
			sub, _ := v1.(map[string]any)
			for k2, v2 := range sub {
				configViper.Set(k1+"."+k2, v2)
			}
		} else {
			configViper.Set(k1, v1)
		}
	}
	configViper.AddConfigPath("./config/")
	configViper.SetConfigName("aaa")
	configViper.SetConfigType("yaml")
	if err := configViper.SafeWriteConfig(); err != nil {
		return err
	}
	return nil
}
