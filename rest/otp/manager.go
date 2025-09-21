package otp

import "github.com/go-redis/redis/v8"

type Manager struct{
	RedisClient *redis.Client 
}
func NewManager(redisClient *redis.Client) *Manager{
	return &Manager{RedisClient: redisClient}
}
