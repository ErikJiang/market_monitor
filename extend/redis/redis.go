package redis

import (
	"time"
	"strconv"
	"github.com/gomodule/redigo/redis"
	"github.com/JiangInk/market_monitor/config"
)

var redisConn *redis.Pool

// GetRedisConn 获取 Redis 客户端连接
func GetRedisConn() *redis.Pool {
	return redisConn
}

// Setup 创建 Redis 连接
func Setup() error {
	redisConn = &redis.Pool{
		MaxIdle: config.RedisConf.MaxIdle,
		MaxActive: config.RedisConf.MaxActive,
		IdleTimeout: config.RedisConf.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisConf.Host+":"+strconv.Itoa(config.RedisConf.Port))
			if err != nil {
				return nil, err
			}
			if config.RedisConf.Password != "" {
				if _, err := c.Do("AUTH", config.RedisConf.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},

	}
	return nil
}


