package redis

import (
	"context"
	"nfldyprdn/maupesen/internal/config"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client
var Ctx = context.Background()

func InitClient() {
	Client = redis.NewClient(&redis.Options{
		Addr: config.C.Redis.Addr,
	})
	if err := Client.Ping(Ctx).Err(); err != nil {
		panic("Redis connection failed: " + err.Error())
	}
}
