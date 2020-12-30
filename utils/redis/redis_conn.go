package redisutils

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/roharon/rpdly-go-url/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ctx = context.Background()

type Redis struct {
	rdb redis.Client
}

func RedisClient() Redis {
	configuration := config.GetConfig()

	password := configuration.REDIS_PASSWORD

	options := redis.Options{
		Addr: configuration.REDIS_ADDRESS,
		DB:   0,
	}

	if password != "" {
		options.Password = password
	}

	rdb := redis.NewClient(&options)

	return Redis{
		rdb: *rdb,
	}
}

func (rds *Redis) Get(key string) (string, error) {
	val, err := rds.rdb.Get(ctx, key).Result()

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

func (rds *Redis) Set(key string, value string) error {
	err := rds.rdb.Set(ctx, key, value, 0).Err()

	if err != nil {
		return err
	} else {
		return nil
	}
}
