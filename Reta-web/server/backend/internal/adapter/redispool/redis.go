package redispool

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type Options struct {
	Addr     string
	Username string
	Password string
	DB       int
}

func NewRedisOptions(conf *viper.Viper) Options {
	return Options{
		Addr:     conf.GetString("redis.host"),
		Username: conf.GetString("redis.user"),
		Password: conf.GetString("redis.password"),
		DB:       conf.GetInt("redis.database"),
	}
}

func NewRedisPool(redisOptions Options) *redis.Client {
	redisPool := redis.NewClient(&redis.Options{
		Addr:     redisOptions.Addr,
		Username: redisOptions.Username,
		Password: redisOptions.Password,
		DB:       redisOptions.DB,
	})
	// ! resullt
	result, err := redisPool.Ping(context.Background()).Result()
	if err != nil {
		return nil
	}

	return redisPool
}
