package config

import (
	"github.com/go-redis/redis"
	"time"
)

type Config struct {
	ServerRole string
	ServerID   string

	redis *redis.Client
}

var Conn = &Config{}

func (c *Config) Redis() *redis.Client {
	if c.redis == nil {
		redisConf := &redis.Options{
			Addr:        ":6379",
			DB:          0,
			MaxRetries:  2,
			IdleTimeout: time.Minute * 5,
		}

		c.redis = redis.NewClient(redisConf)
	}

	return c.redis
}
