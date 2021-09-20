package main

import (
	"encoding/json"
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
	successSismember(client)
}

func simpleCache(c *redis.Client) {
	key := "foo"
	err := c.Set(key, "bar", 0).Err()
	if err != nil {
		log.Println(err)
	}

	val, err := c.Get(key).Result()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(val)
}

func successSismember(c *redis.Client) {
	key := "foofoo"
	jsonStr := `{"filename": "r7iPc2eE4KNzfXHaKJUkbfPGoXaruE.json", "user": "77xxx"}`
	json_, err := json.Marshal(jsonStr)
	if err != nil {
		log.Println(err)
	}
	c.SAdd(key, jsonStr, json_)

	fmt.Println(c.SIsMember(key, jsonStr))
	fmt.Println(c.SIsMember(key, json_))
}
