package redis_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golearnku/sdk-provider/redis"
)

func ExampleNewClient() {
	var err error
	config := redis.Config{
		Addr:     "127.0.0.1:6379",
		Password: "",
		PoolSize: 100,
		DB:       0,
	}
	client, err = redis.NewClient(config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("client: %+v \n", client)
}

func ExampleGetConnDB() {
	db := redis.GetConnDB()
	if err := db.Set(context.Background(), "lock", "1", time.Second*2).Err(); err != nil {
		log.Fatal(err)
	}

	val := db.Get(context.Background(), "lock").Val()
	fmt.Printf("val: %s \n", val)
	// OutPut:val: 1
}

func ExampleRedis_Ping() {
	if err := client.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	// OutPut:
}

func ExampleRedis_Close() {
	if err := client.Close(); err != nil {
		log.Fatal(err)
	}
	// OutPut:
}
