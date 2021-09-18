package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("aaa")

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	simpleCache(client)
}

func simpleCache(c *redis.Client) {
	key := "foo"
	ctx := context.Background()
	err := c.Set(ctx, key, "bar", 0).Err()
	if err != nil {
		log.Println(err)
	}

	val, err := c.Get(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(val)
}
