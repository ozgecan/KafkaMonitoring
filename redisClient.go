package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func redisSet(client *redis.Client, key string, value string) {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func redisGet(client *redis.Client, key string) string {
	value, err := client.Get(key).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
	return value
}
