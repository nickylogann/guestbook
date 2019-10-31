package redis

import (
	"github.com/gomodule/redigo/redis"

	"github.com/nickylogan/guestbook/internal/app/repository/visitor"
)

type redisRepository struct {
	client *redis.Pool
	key    string
}

func NewRedisRepository(client *redis.Pool, key string) visitor.Repository {
	return &redisRepository{
		client: client,
		key:    key,
	}
}
