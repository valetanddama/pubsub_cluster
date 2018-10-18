package services

import (
	"github.com/valetanddama/cluster/pkg/config"
	"log"
)

func GetListOfWrongMessages() {
	result, err := config.Conn.Redis().SMembers(config.KeyErrors).Result()

	if err != nil {
		log.Println(err)
		return
	}

	for _, errorMessage := range result {
		log.Println("Error message: " + errorMessage)

		if err := config.Conn.Redis().SRem(config.KeyErrors, errorMessage).Err(); err != nil {
			log.Println(err)
		}
	}
}
