package cache

import (
	"BlogProject/logger"
	"github.com/go-redis/redis"
	"strconv"
)

//RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// InitRedis 初始化redis链接
func InitRedis(RedisAddr, RedisPw, RedisDbName string) {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logger.Logger.Info(err)
		panic(err)
	}
	RedisClient = client
}
