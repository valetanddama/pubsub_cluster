package services

import (
	"github.com/go-redis/redis"
	"github.com/valetanddama/cluster/pkg/config"
	"log"
	"time"
)

func CheckServerRole() {
	serverID, err := config.Conn.Redis().Get(config.KeyPublisher).Result()

	if err != nil && err != redis.Nil {
		log.Println(err)
		return
	}

	switch serverID {
	case config.Conn.ServerID:
		if err := config.Conn.Redis().Expire(config.KeyPublisher, 2*time.Second).Err(); err != nil {
			log.Println(err)
			return
		}

		config.Conn.ServerRole = config.RoleGeneratorMessages
	default:
		config.Conn.ServerRole = config.RoleHandlerMessages

		if result, err := config.Conn.Redis().SetNX(config.KeyPublisher, config.Conn.ServerID, 2*time.Second).Result(); err != nil {
			log.Println(err)
		} else if result == true {
			config.Conn.ServerRole = config.RoleGeneratorMessages
		}
	}
}
