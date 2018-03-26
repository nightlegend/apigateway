package redis

import (
	"fmt"

	"github.com/go-redis/redis"
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
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return client
}
