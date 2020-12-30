package coverage

import (
	"os"
	"testing"

	"github.com/alicebob/miniredis"
	redis "github.com/roharon/rpdly-go-url/utils/redis"
)

func TestRedisClient(t *testing.T) {
	mr, err := miniredis.Run()
	os.Setenv("REDIS_PASSWORD", "")
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection",
			err)
	}
	os.Setenv("REDIS_ADDRESS", mr.Addr())
	// set miniredis's address for testing

	rds := redis.RedisClient()

	testKey := "test:value"
	testVal := "3"
	err = rds.Set(testKey, testVal)
	if err != nil {
		t.Error("redis: Set function error", err)
	}

	val, err := rds.Get(testKey)
	if err != nil {
		t.Error("redis: value not exist", err)
	}

	if testVal != val {
		t.Error("redis: test value is not same val that return")
	} else {
		t.Logf("real value: %s || in redis value: %s", testVal, val)
	}
}
