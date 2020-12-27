package main

import (
	"github.com/alicebob/miniredis"
	"github.com/roharon/rpdly-go-url/utils"
	"testing"
)

func TestRedisClient(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection",
			err)
	}

	client := redisutils.RedisClient(mr.Addr(), "")

	testKey := "test:value"
	testVal := "3"
	err = redisutils.Set(client, testKey, testVal)
	if err != nil {
		t.Error("redis: Set function error", err)
	}

	val, err := redisutils.Get(client, testKey)
	if err != nil {
		t.Error("redis: value not exist", err)
	}

	if testVal != val {
		t.Error("redis: test value is not same val that return")
	} else {
		t.Logf("real value: %s || in redis value: %s", testVal, val)
	}
}
