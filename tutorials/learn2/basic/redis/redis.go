package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"micro_learn/tutorials/learn2/basic/config"
	"sync"
)


var once sync.Once
var redisCliet *redis.Client



func Init() {
	once.Do(func() {
		initRedisClient()
	})
}


func initRedisClient() {
	redisConfig := config.GetRedisConf()

	if redisConfig == nil {
		panic("read redis config failed")
	}

	if !redisConfig.Enabled {
		return
	}

	redisCliet = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Conn,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DbNum,    // use default DB
	})

	pong, err := redisCliet.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}

	fmt.Println(pong)
}

func GetRedisClient() *redis.Client {
	once.Do(func() {
		initRedisClient()
	})

	return redisCliet
}