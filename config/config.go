package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type config struct {
	SERVER_ADDRESS string `env:"SERVER_ADDRESS" envDefault:"0.0.0.0:3000"`
	REDIS_ADDRESS  string `env:"REDIS_ADDRESS"`
	REDIS_PASSWORD string `env:"REDIS_PASSWORD"`
	PROXY_PORT     string `env:"PROXY_PORT" envDefault:":8081"`
}

var (
	ENV_TEST string = getBasePath() + "/../.env.test"
	ENV_PROD string = getBasePath() + "/../.env"
)

func getBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	return basePath
}

func init() {
	currentEnv := os.Getenv("ENVIRONMENT")

	if currentEnv == "prod" {
		if err := godotenv.Load(ENV_PROD); err != nil {
			log.Fatalf("Problem loading %s file: %v", ENV_PROD, err)
		}
		log.Printf("==PROD MODE: %s==", currentEnv)
	} else {
		log.Printf("==TEST MODE %s==", currentEnv)
		if currentEnv != "test" {
			log.Print("You can set ENVIRONMENT prod | test")
		}

		if err := godotenv.Load(ENV_TEST); err != nil {
			log.Fatalf("Problem loading %s file: %v", ENV_TEST, err)
		}
	}
}

func GetConfig() config {
	configuration := config{}

	if err := env.Parse(&configuration); err != nil {
		log.Printf("%+v", err)
	}
	return configuration
}
