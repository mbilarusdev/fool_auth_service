package infrasturcture

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

func PingRedis() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	password := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	ctx := context.Background()

	pingCmd := rdb.Ping(ctx)

	LogInfo(fmt.Sprintf("From Redis: %v", pingCmd.Val()))

	return rdb
}
