package cache

import (
    "github.com/go-redis/redis/v8"
    "context"
)

var ctx = context.Background()

func ConnectRedis() *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
}
