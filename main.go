package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func main() {
	ctx := context.Background()
	keyChan := make(chan string)
	valueChan := make(chan string)

	go produce(ctx)
	go func() {
		consume(ctx, keyChan, valueChan)
	}()

	client := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "redisPassword", DB: 10})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)
	for {
		fmt.Println("come to for loop")
		select {
		case <-keyChan:
			if <-valueChan != "" {
				fmt.Println("valueChan is not nil")
				redisSet(client, <-keyChan, <-valueChan)
			}
			fmt.Println("valueChan is  = ", redisGet(client, <-keyChan))

		}
	}

}
