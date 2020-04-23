package config

import "fmt"

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

var RedisConfigSetting *RedisConfig

func init() {
	redisHost, err := LoadConfig("REDIS_HOST")
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	redisPwd, err := LoadConfig("REDIS_PASSWORD")
	if err != nil {
		fmt.Printf("err:%v", err)
	}

	RedisConfigSetting = &RedisConfig{
		Host:     redisHost,
		Port:     "6379",
		Password: redisPwd,
	}
}
