package redis

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

func (r redisRepository) Increment(ctx context.Context) (res int, err error) {
	conn := r.client.Get()
	defer conn.Close()

	res, err = redis.Int(conn.Do("INCR", r.key))
	return
}
