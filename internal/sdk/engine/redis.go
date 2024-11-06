package engine

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/weapon-team/weapon/internal/sdk/runtime"
)

func initRedisClient() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:           runtime.Setting.Redis.Addr,
		Username:       runtime.Setting.Redis.Username,
		Password:       runtime.Setting.Redis.Password,
		MaxActiveConns: runtime.Setting.Redis.MaxActive,
	})
	err := rdb.Ping(context.Background()).Err()
	return rdb, err
}
