package dao

import (
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	// 初始化redis客户端连接
	redisCli *redis.Client
)

// init 初始化redis连接
func init() {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
	})
	// 检查客户端连接是否正确
	_, err := redisCli.Ping(ctx).Result()
	if err != nil {
		log.Fatalln(err)
	}
}

