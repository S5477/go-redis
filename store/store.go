package store

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const ADDR = "localhost:6379"

var ctx = context.Background()
var rbase = redis.NewClient(&redis.Options{
	Addr:     ADDR,
	Password: "",
	DB:       0,
})

func SetString(key string, value string) {

	err := rbase.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func GetString(key string) string {

	val, err := rbase.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return val
}
