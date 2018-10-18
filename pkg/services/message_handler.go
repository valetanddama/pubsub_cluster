package services

import (
	"github.com/go-redis/redis"
	"github.com/valetanddama/cluster/pkg/config"
	"log"
	"math/rand"
	"time"
)

func GetMessage() {
	if config.Conn.ServerRole == config.RoleHandlerMessages {
		message, err := config.Conn.Redis().BLPop(time.Second, config.KeyMessages).Result()

		if err != nil && err != redis.Nil {
			log.Println(err)
			return
		}

		if len(message) == 2 {
			if rand.Intn(100) < 5 {
				config.Conn.Redis().SAdd(config.KeyErrors, message[1])
			} else {
				log.Println("Message: " + message[1])
			}
		}
	}

}

func SendMessage() {
	if config.Conn.ServerRole == config.RoleGeneratorMessages {
		bytes := make([]byte, 20)

		for i := 0; i < cap(bytes); i++ {
			bytes[i] = byte(65 + rand.Intn(25))
		}

		if err := config.Conn.Redis().RPush(config.KeyMessages, string(bytes)).Err(); err != nil {
			log.Println(err)
		}
	}
}
