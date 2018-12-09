package model

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Example is a simple connection example.
func Example() {
	client := redis.NewClient(&redis.Options{
		Addr:     "store:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
