package services

import (
	"github.com/valetanddama/cluster/pkg/config"
	"log"
	"time"
)

func CheckServerRole() {
	result, err := config.Conn.Redis().SetNX(config.KeyPublisher, config.Conn.ServerID, 2*time.Second).Result()

	if err != nil {
		log.Println(err)
		return
	}

	if result == false {
		serverID, err := config.Conn.Redis().Get(config.KeyPublisher).Result()

		if err != nil {
			log.Println(err)
			return
		}

		if serverID == config.Conn.ServerID {
			if err := config.Conn.Redis().Expire(config.KeyPublisher, 2*time.Second).Err(); err != nil {
				log.Println(err)
				return
			}

			config.Conn.ServerRole = config.RoleGeneratorMessages
		} else {
			config.Conn.ServerRole = config.RoleHandlerMessages
		}
	} else {
		config.Conn.ServerRole = config.RoleGeneratorMessages
	}
}
