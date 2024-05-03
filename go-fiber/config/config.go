package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvValueFromKey(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Error Occured When Loading .env files %v", err.Error())
	}
	return os.Getenv(key)
}
