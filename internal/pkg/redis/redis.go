package redis

import (
	"fmt"

	"github.com/Ted-bug/open-api/config"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedis() error {
	option := GetRedisConf()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", option.Host, option.Port),
		Password: option.Password,
		DB:       option.Db,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println("redisClient init success")
	return nil
}

// 获取redis的配置信息；
// 如果是连接到集群的话，请修改配置项
func GetRedisConf() config.Redis {
	return config.AppConfig.Redis
}

func CloseRedis() {
	fmt.Println("close redis...")
	RedisClient.Close()
}
