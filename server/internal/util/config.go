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
	config.ServerAddress = os.Getenv("SERVER_ADDRESS")
	config.ClerkSK = os.Getenv("CLERK_SK")
	config.RedisAddress = os.Getenv("REDIS_ADDRESS")
	return
}
