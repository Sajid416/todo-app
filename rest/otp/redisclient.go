package otp

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Ctx=context.Background()

func NewClient(addr string) *redis.Client{
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}