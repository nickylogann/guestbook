package redis

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

func (r redisRepository) Get(ctx context.Context) (res int, err error) {
	conn := r.client.Get()
	defer conn.Close()

	res, err = redis.Int(conn.Do("GET", r.key))
	return
}
