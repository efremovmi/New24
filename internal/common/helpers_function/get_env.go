package helpers_function

import (
	"News24/internal/models"

	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnvParams() models.Config {
	err := godotenv.Load(`.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return models.Config{
		ADR_AUTH:                  os.Getenv("ADR_AUTH"),
		PROTOCOL_WITH_DOMAIN_AUTH: os.Getenv("PROTOCOL_WITH_DOMAIN_AUTH"),

		ADR_CONTROL_USERS:                  os.Getenv("ADR_CONTROL_USERS"),
		PROTOCOL_WITH_DOMAIN_CONTROL_USERS: os.Getenv("PROTOCOL_WITH_DOMAIN_CONTROL_USERS"),

		ADR_NEWS:                  os.Getenv("ADR_NEWS"),
		PROTOCOL_WITH_DOMAIN_NEWS: os.Getenv("PROTOCOL_WITH_DOMAIN_NEWS"),

		POSTGRES_HOST:        os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:        os.Getenv("POSTGRES_PORT"),
		POSTGRES_USER:        os.Getenv("POSTGRES_USER"),
		POSTGRES_BD_NAME:     os.Getenv("POSTGRES_BD_NAME"),
		POSTGRES_PASSWORD:    os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_TABLE_USERS: os.Getenv("POSTGRES_TABLE_USERS"),
		HASH_SALT:            os.Getenv("HASH_SALT"),

		POSTGRES_TABLE_NEWS: os.Getenv("POSTGRES_TABLE_NEWS"),
	}
}
