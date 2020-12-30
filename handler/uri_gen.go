package handler

import (
	"context"
	"strconv"

	randomutils "github.com/roharon/rpdly-go-url/utils/random"
	redis "github.com/roharon/rpdly-go-url/utils/redis"

	pb "github.com/roharon/rpdly-go-url/protobuf/uri/v1"
)

type RouteServer struct{}

const RANDOM_LENGTH = "URL:LENGTH"
const URL_PREFIX = "URL:content:"
const DEFAULT_LENGTH = 4

func (s *RouteServer) GetUri(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	key := URL_PREFIX + req.GetUri()

	rdb := redis.RedisClient()
	val, err := rdb.Get(key)

	if err != nil && val == "" {
		return &pb.Response{}, err
	}
	return &pb.Response{Uri: val}, nil
}

func (s *RouteServer) SetUri(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	originUrl := req.GetUri()
	rds := redis.RedisClient()

	val, err := rds.Get(RANDOM_LENGTH)
	if err != nil {
		_ = rds.Set(RANDOM_LENGTH, strconv.Itoa(DEFAULT_LENGTH))
	}

	valInt, errValInt := strconv.Atoi(val)
	if errValInt != nil {
		valInt = DEFAULT_LENGTH
	}

	shortUrl := randomutils.RandomString(valInt)

	err = rds.Set(URL_PREFIX+shortUrl, originUrl)
	if err != nil {
		return &pb.Response{}, err
	}

	return &pb.Response{Uri: shortUrl}, nil
}
