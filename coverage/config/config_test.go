package coverage

import (
	"os"
	"testing"

	"github.com/roharon/rpdly-go-url/config"
)

func TestGetConfig(t *testing.T) {

	conf := config.GetConfig()

	if conf.REDIS_ADDRESS != os.Getenv("REDIS_ADDRESS") {

		t.Error("Not correct", conf.REDIS_ADDRESS)
	}

	if conf.SERVER_ADDRESS != os.Getenv("SERVER_ADDRESS") {
		t.Error("Not correct", conf.REDIS_ADDRESS)
	}

	if conf.REDIS_PASSWORD != os.Getenv("REDIS_PASSWORD") {
		t.Error("Not correct", conf.REDIS_ADDRESS)
	}
}
