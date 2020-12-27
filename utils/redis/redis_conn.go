package redisutils

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ctx = context.Background()

func RedisClient(addr string, password string) *redis.Client {
	if addr == "" && password == "" {
		_ = godotenv.Load("../.env")
		addr = os.Getenv("redis_address")
		password = os.Getenv("redis_password")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return rdb
}

func Get(rdb *redis.Client, key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		log.Printf("%s does not exist", key)
		err := status.Errorf(codes.NotFound, "%s does not exist", key)
		return "", err
	} else if err != nil {
		// error occurs
		return "", err
	}

	return val, nil
}

func Set(rdb *redis.Client, key string, value string) error {
	err := rdb.Set(ctx, key, value, 0).Err()

	if err != nil {
		return err
	} else {
		return nil
	}
}
