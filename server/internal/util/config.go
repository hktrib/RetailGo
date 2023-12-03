package util

import (
	"fmt"
	"os"
)

type Config struct {
	DB_DRIVER            string // `"DB_DRIVER"`
	DB_SOURCE            string // "DB_SOURCE"`
	SERVER_ADDRESS       string // `"SERVER_ADDRESS"`
	CLERK_SK             string // `"CLERK_SK"`
	CLERK_WEBHOOK_SECRET string // `"CLERK_WEBHOOK_SECRET"`
	REDIS_HOSTNAME       string // `"REDIS_HOSTNAME"`
	REDIS_PORT           string // `"REDIS_PORT"`
	REDIS_PASSWORD       string // `"REDIS_PASSWORD"`
	STRIPE_SK            string // `"STRIPE_SK"`
	HOST                 string // `"HOST"`
	WEAVIATE_HOSTNAME    string // `"WEAVIATE_HOSTNAME"`
	WEAVIATE_SK          string // `"WEAVIATE_SK"`

	STRIPE_WEBHOOK_SECRET string //
}

func LoadConfig() (config Config, err error) {
	// viper.AddConfigPath("./util/config")
	// viper.SetConfigName(".env")
	// viper.SetConfigType("env")

	// viper.AutomaticEnv()

	config.DB_DRIVER = os.Getenv("DB_DRIVER")
	config.DB_SOURCE = os.Getenv("DB_SOURCE")
	fmt.Println(config.DB_SOURCE)
	config.SERVER_ADDRESS = envPortOr("8080")
	config.CLERK_SK = os.Getenv("CLERK_SK")
	config.CLERK_WEBHOOK_SECRET = os.Getenv("CLERK_WEBHOOK_SECRET")
	config.REDIS_HOSTNAME = os.Getenv("REDIS_HOSTNAME")
	config.REDIS_PORT = os.Getenv("REDIS_PORT")
	config.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	config.STRIPE_SK = os.Getenv("STRIPE_SK")
	config.WEAVIATE_HOSTNAME = os.Getenv("WEAVIATE_HOSTNAME")
	config.WEAVIATE_SK = os.Getenv("WEAVIATE_SK")
	config.STRIPE_WEBHOOK_SECRET = os.Getenv("STRIPE_WEBHOOK_SECRET")

	// fmt.Println(config)

	return
}

func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return port
}
