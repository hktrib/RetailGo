package util

import (
	"os"
)

type Config struct {
	DBDriver      string // `"DB_DRIVER"`
	DBSource      string // "DB_SOURCE"`
	ServerAddress string // `"SERVER_ADDRESS"`
	ClerkSK       string // `"CLERK_SK"`
	RedisAddress  string // `"REDIS_ADDRESS"`
}

func LoadConfig() (config Config, err error) {
	// viper.AddConfigPath("./util/config")
	// viper.SetConfigName(".env")
	// viper.SetConfigType("env")

	// viper.AutomaticEnv()

	config.DBDriver = os.Getenv("DB_DRIVER")
	config.DBSource = os.Getenv("DB_SOURCE")
	config.ServerAddress = envPortOr("3000")
	config.ClerkSK = os.Getenv("CLERK_SK")
	config.RedisAddress = os.Getenv("REDIS_ADDRESS")
	return
}

func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("SERVER_ADDRESS"); envPort != "" {
		return envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return port
}
