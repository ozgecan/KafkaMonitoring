package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	topic          = "foo"
	broker1Address = "localhost:9092"
)

func produce(ctx context.Context) {
	i := 0
	w := kafka.NewWriter(kafka.WriterConfig{Brokers: []string{broker1Address},
		Topic: topic})

	for {
		err := w.WriteMessages(ctx, kafka.Message{Key: []byte(strconv.Itoa(i)), Value: []byte("this is message" + strconv.Itoa(i))})
		if err != nil {
			panic("could not write message" + err.Error())
		}
		fmt.Println("writes:", i)
		i++
		time.Sleep(time.Second)
	}
}

func consume(ctx context.Context) {
	l := log.New(os.Stdout, "kafka reader: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
		Logger:  l,
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message" + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
	}
}

func main() {
	ctx := context.Background()
	go produce(ctx)
	consume(ctx)
}
