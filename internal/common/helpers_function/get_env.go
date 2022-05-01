package helpers_function

import (
	"News24/internal/models"

	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnvParams() models.Config {
	err := godotenv.Load(`/home/max/News24/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return models.Config{
		ADDR_AUTH:                 os.Getenv("ADDR_AUTH"),
		PROTOCOL_WITH_DOMAIN_AUTH: os.Getenv("PROTOCOL_WITH_DOMAIN_AUTH"),

		ADDR_CONTROL_USERS:                 os.Getenv("ADDR_CONTROL_USERS"),
		PROTOCOL_WITH_DOMAIN_CONTROL_USERS: os.Getenv("PROTOCOL_WITH_DOMAIN_CONTROL_USERS"),

		POSTGRES_HOST:        os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:        os.Getenv("POSTGRES_PORT"),
		POSTGRES_USER:        os.Getenv("POSTGRES_USER"),
		POSTGRES_BD_NAME:     os.Getenv("POSTGRES_BD_NAME"),
		POSTGRES_PASSWORD:    os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_TABLE_USERS: os.Getenv("POSTGRES_TABLE_USERS"),
		HASH_SALT:            os.Getenv("HASH_SALT"),
	}
}
