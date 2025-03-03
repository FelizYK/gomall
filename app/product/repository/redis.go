package repository

import (
	"context"

	"github.com/FelizYK/gomall/app/product/conf"
	redis "github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
)

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
