package impl

import (
	"canbefree/go-course/infrastructure/vars"

	"github.com/gomodule/redigo/redis"
)

type RedisCache struct {
}

func NewRedisCache() *RedisCache {
	return &RedisCache{}
}

func (cache *RedisCache) Get(key string) (string, error) {
	conn := vars.RedisPool.Get()
	defer conn.Close()
	return redis.String(conn.Do("get", key))
}

func (cache *RedisCache) Set(key, val string) error {
	conn := vars.RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("setex", key, 60, val)
	if err != nil {
		return err
	}
	return nil
}
