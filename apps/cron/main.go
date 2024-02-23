package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	goredis "github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	redis := goredis.NewClient(&goredis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PWD"),
		DB:       0,
	})

	// Publish
	err := redis.Publish(ctx, "triger", "fire").Err()
	if err != nil {
		panic(err)
	}

	// Subscribe
	pubsub := redis.Subscribe(ctx, "triger")
	defer pubsub.Close()

	sub := pubsub.Channel()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	fmt.Println("Cron job start :", time.Now())

	for {
		select {
		case <-ticker.C:
			err := doGet()
			if err != nil {
				println("Error:", err.Error())
			}
		case msg := <-sub:
			if msg.Payload == "fire" {
				fmt.Println("Cron job finish :", time.Now())
				return
			}
		}
	}
}

func doGet() error {
	host := os.Getenv("TARGET_HOST")

	client := resty.New()
	resp, err := client.R().EnableTrace().Get("http://" + host)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return err
	}

	return nil
}
