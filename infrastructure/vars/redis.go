package vars

import (
	"canbefree/go-course/infrastructure/config"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func GetConn() (redis.Conn, error) {
	c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", config.RedisConfigSetting.Host, config.RedisConfigSetting.Port))
	if err != nil {
		fmt.Printf("err:%err", err)
		return nil, err
	}
	if config.RedisConfigSetting.Password != "" {
		if _, err := c.Do("auth", config.RedisConfigSetting.Password); err != nil {
			c.Close()
			return nil, err
		}
	}
	return c, nil
}

func CreateRedisPool() (*redis.Pool, error) {
	return &redis.Pool{
		MaxIdle:   10,
		MaxActive: 5,
		Dial: func() (redis.Conn, error) {
			return GetConn()
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, nil
}

var RedisPool *redis.Pool

func init() {
	var err error
	RedisPool, err = CreateRedisPool()
	if err != nil {
		fmt.Println("err:%v", err)
	}
	return
}
