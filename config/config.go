package config

import (
	_ "embed"
	"errors"
	"fmt"
	"os"

	"github.com/Ted-bug/open-api/internal/constants"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Mode   string `mapstructure:"mode"`
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	Gorm   Gorm   `mapstructure:"gorm"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
	Logger Logger `mapstructure:"logger"`
}

type Gorm struct {
	LogLevel string `mapstructure:"log_level"`
}

type Mysql struct {
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DbName          string `mapstructure:"db_name"`
	Charset         string `mapstructure:"charset"`
	Prefix          string `mapstructure:"prefix"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxIdleTime int    `mapstructure:"conn_max_idle_time"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	Prefix   string `mapstructure:"prefix"`
}

type Logger struct {
	Package    string   `mapstructure:"package"`
	List       []string `mapstructure:"list"`
	Level      string   `mapstructure:"level"`
	Output     string   `mapstructure:"output"`
	Path       string   `mapstructure:"path"`
	Filename   string   `mapstructure:"filename"`
	MaxSize    int      `mapstructure:"max_size"`
	MaxAge     int      `mapstructure:"max_age"`
	MaxBackups int      `mapstructure:"max_backups"`
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
	fmt.Println("config init success")
	return nil
}

//go:embed config_example.yaml
var configExample string

// 创建配置文件示例
func CreateConfigFile(filename string) error {
	if filename == "" {
		filename = "config"
	}
	path := "./config/" + filename + ".yaml"
	if _, err := os.Stat(path); err == nil || !os.IsNotExist(err) {
		return errors.New("the file is exist: " + path)
	}
	if err := os.WriteFile(path, []byte(configExample), 0644); err != nil {
		return err
	}
	return nil
}
