package redis_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/bsm/redislock"
	"github.com/golearnku/sdk-provider/redis"
)

var client *redis.Redis

func TestMain(t *testing.M) {
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
	t.Run()
}

func TestNewClient(t *testing.T) {
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
	t.Logf("client: %+v \n", client)
}

func TestGetConnDB(t *testing.T) {
	db := redis.GetConnDB()
	if err := db.Set(context.Background(), "lock", "1", time.Second*2).Err(); err != nil {
		t.Error(err)
	}

	val := db.Get(context.Background(), "lock").Val()
	t.Logf("val: %s \n", val)
}

func TestRedis_Ping(t *testing.T) {
	if err := client.Ping(context.Background()); err != nil {
		t.Error(err)
	}
}

func TestRedis_Close(t *testing.T) {
	if err := client.Close(); err != nil {
		t.Error(err)
	}
}

func TestRedisLock(t *testing.T) {
	db := GetConnDB()
	locker := redislock.New(db)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		_, err := locker.Obtain("redis-lock", 100*time.Millisecond, nil)
		if err == redislock.ErrNotObtained {
			fmt.Println("无法获得锁")
			continue
		} else if err != nil {
			log.Fatalln(err)
		}

		// Don't forget to defer Release.
		//defer lock.Release()
		fmt.Println("I have a lock!")
	}
}
