package redis

import (
	"strconv"
	"github.com/go-redis/redis"
	"github.com/JiangInk/market_monitor/config"
)

var redisCli *redis.Client

// GetRedisClient 获取 Redis 客户端连接实例
func GetRedisClient() *redis.Client {
	return redisCli
}

// Setup 创建 Redis 连接
func Setup() {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisConf.Host + ":"+strconv.Itoa(config.RedisConf.Port),
		Password: config.RedisConf.Password,
		DB: config.RedisConf.DBNum,
	})
	redisCli = client
}
