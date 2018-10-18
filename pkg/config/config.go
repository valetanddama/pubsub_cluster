package config

import "github.com/go-redis/redis"

type Config struct {
	ServerRole string
	ServerID   string

	redis *redis.Client
}

var Conn = &Config{}

func (c *Config) Redis() *redis.Client {
	if c.redis == nil {
		redisConf := &redis.Options{
			Addr: ":6379",
			DB:   0,
		}

		c.redis = redis.NewClient(redisConf)
	}

	return c.redis
}

func (c *Config) Ping() {
	if err := c.Redis().Ping().Err(); err != nil {
		c.redis = nil
	}
}
