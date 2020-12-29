package coverage

import (
	"math/rand"
	"testing"
	"time"

	randomutils "github.com/roharon/rpdly-go-url/utils/random"
)

func TestRandomString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	size := 6
	result := randomutils.RandomString(size)

	if len(result) != 6 {
		t.Error("string length is incorrect")
	}
	t.Logf("generate result : %s", result)

	size = 0
	result = randomutils.RandomString(size)

	if result != "" {
		t.Error("random string incorrect when length 0")
	}
	t.Logf("random string len 0 : %s", result)
}
