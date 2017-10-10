package cache

import (
	"github.com/nightlegend/apigateway/core/utils/redis"
	"log"
)

func SaveLoginSession(loginSession string) {
	client := redis.NewClient()
	err := client.Set("sessionId", "hgdes=sdsa=dasje23", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("sessionId").Result()
	if err != nil {
		panic(err)
	}
	log.Println("sessionId", val)
}
