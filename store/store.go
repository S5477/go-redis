package store

import (
	"context"

	"github.com/S5477/go-redis/models"
	"github.com/redis/go-redis/v9"
)

const ADDR = "localhost:6379"

var ctx = context.Background()
var rbase = redis.NewClient(&redis.Options{
	Addr:     ADDR,
	Password: "",
	DB:       0,
})

func SetString(s models.Add) {

	println("start", s.Key, s.Value)
	err := rbase.Set(ctx, s.Key, s.Value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func GetString(key string) string {
	val, _ := rbase.Get(ctx, key).Result()

	return val
}
