package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/nickylogan/guestbook/internal/pkg/utils/config"
)

// NewConnection creates a new redis connection
func NewConnection(cfg config.RedisConfig) *redis.Pool {
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return connectRedis(cfg.URI)
		},
		MaxActive:   cfg.MaxActive,
		IdleTimeout: time.Duration(10) * time.Second,
		Wait:        true,
	}
	return redisPool
}

func connectRedis(uri string) (redis.Conn, error) {
	c, err := redis.Dial("tcp", uri)
	if err != nil {
		return nil, err
	}
	return c, nil
}
