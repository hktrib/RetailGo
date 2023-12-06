package util

import (
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
	WEAVIATE_SK    		 string // `"WEAVIATE_SK"`

	SUPABASE_URL          string // `"SUPABASE_URL"`
	SUPABASE_KEY          string // `"SUPABASE_KEY"`

	STRIPE_WEBHOOK_SECRET string //
}

func LoadConfig() (config Config, err error) {

	// Supabase
	config.SUPABASE_URL = os.Getenv("SUPABASE_URL")
	config.SUPABASE_KEY = os.Getenv("SUPABASE_KEY")
	
	// PostgresDB
	config.DB_DRIVER = os.Getenv("DB_DRIVER")
	config.DB_SOURCE = os.Getenv("DB_SOURCE")
	
	// Clerk
	config.CLERK_SK = os.Getenv("CLERK_SK")
	config.CLERK_WEBHOOK_SECRET = os.Getenv("CLERK_WEBHOOK_SECRET")
	
	// Redis
	config.REDIS_HOSTNAME = os.Getenv("REDIS_HOSTNAME")
	config.REDIS_PORT = os.Getenv("REDIS_PORT")
	config.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	
	// Weaviate
	config.WEAVIATE_HOSTNAME = os.Getenv("WEAVIATE_HOSTNAME")
	config.WEAVIATE_SK = os.Getenv("WEAVIATE_SK")
	
	// Stripe
	config.STRIPE_SK = os.Getenv("STRIPE_SK")
	config.STRIPE_WEBHOOK_SECRET = os.Getenv("STRIPE_WEBHOOK_SECRET")
	
	// Server Address
	config.SERVER_ADDRESS = envPortOr("8080")
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
