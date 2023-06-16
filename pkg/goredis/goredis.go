package goredis

import (
	"admin-server/pkg/logger"
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"log"
)

var rdb *redis.Client

func InitRedis() *redis.Client {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	db := viper.GetInt("redis.db")
	password := viper.GetString("redis.password")
	redis.SetLogger(logger.NewMyRedisLogger())
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("failed to connect redis,err:" + err.Error())
	}
	log.Println("Redis Client: " + pong)
	return rdb
}

func GetRedisDB() *redis.Client {
	return rdb
}

func GetRedisDBWithContext(c context.Context) *redis.Client {
	redis.SetLogger(logger.NewMyRedisLoggerWithContext(c))
	return rdb
}
