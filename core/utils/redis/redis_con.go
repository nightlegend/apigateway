package redis

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// NewClient :
// new connect object for redis.
func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	log.Info(pong, err)
	// Output: PONG <nil>
	return client
}
