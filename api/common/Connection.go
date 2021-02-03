package common

import (
	"context"

	"github.com/go-redis/redis/v8")


type Redis struct {
	Client *redis.Client
}

var (
	RedisAddress = "localhost:6379"
	ctx = context.TODO()
)

func Connection(address string) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddress,
		Password: "",
		DB: 0,
	})

	err := client.Ping(ctx).Err()

	if err != nil {
		return nil, err
	}

	return &Redis{
		Client: client,
	}, nil
}