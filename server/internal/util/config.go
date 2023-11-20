package util

import (
	"fmt"
	"os"
)

type Config struct {
	DBDriver      string // `"DB_DRIVER"`
	DBSource      string // "DB_SOURCE"`
	ServerAddress string // `"SERVER_ADDRESS"`
	ClerkSK       string // `"CLERK_SK"`
	RedisAddress  string // `"REDIS_ADDRESS"`
	StripeSK      string // `"STRIPE_SK"`
}

func LoadConfig() (config Config, err error) {
	// viper.AddConfigPath("./util/config")
	// viper.SetConfigName(".env")
	// viper.SetConfigType("env")

	// viper.AutomaticEnv()

	config.DBDriver = "postgres"                                                                                 //os.Getenv("DB_DRIVER")
	config.DBSource = "postgresql://postgres:76ashcuCoOhkEhgb@db.zvevvgcnviqxagbysekg.supabase.co:5432/postgres" // os.Getenv("DB_SOURCE")
	config.ServerAddress = envPortOr("8080")
	config.ClerkSK = "sk_test_wCmeudOz44ArIXVFbzTjFOOqhbPquW94kdazRMmjfQ"
	fmt.Println("ClerkSK:", config.ClerkSK)
	config.RedisAddress = "redis://default:HBCmEOKMFFGN6525oJombkA6IfnfKaHn@viaduct.proxy.rlwy.net:38806"
	config.StripeSK = "sk_test_51ODz7pHWQUATs9zV4fWYLtRag0GwwLPticrlOe5FqicEWwdnWUlsZkRh90o1YOkt3qsOduJQNSbbUJupkm4i9xLm00hcffWjDm"

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
