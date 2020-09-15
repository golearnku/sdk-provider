package redis

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	internalRedis *Redis
	onceRedis     sync.Once

	Nil         = redis.Nil
	TxFailedErr = redis.TxFailedErr
)

type Redis struct {
	client *redis.Client
}

type Config struct {
	Addr     string //连接地址 127.0.0.1:6379
	Password string //密码
	PoolSize int    //连接池数量
	DB       int    //db select 空间
}

//初始化 Redis
func NewClient(config Config) (*Redis, error) {
	onceRedis.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     config.Addr,
			Password: config.Password,
			PoolSize: config.PoolSize,
			DB:       config.DB,
		})
		internalRedis = &Redis{client: client}
	})
	return internalRedis, nil
}

//获取 Redis 实例
func GetConnDB() *redis.Client {
	if internalRedis == nil || internalRedis.client == nil {
		panic("Redis Client is not initialized")
	}
	return internalRedis.client
}

//redis ping.
func (r *Redis) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

//redis close.
func (r *Redis) Close() error {
	return r.client.Close()
}
