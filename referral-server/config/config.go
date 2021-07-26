package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// const (
// 	DB_HOST              = "127.0.0.1"
// 	DB_USER              = "store-adm"
// 	DB_NAME              = "referral_system"
// 	DB_PASSWORD          = "pass"
// 	API_PORT             = "8080"
// 	SHARED_LINK_ENDPOINT = "http://localhost:8080/acces/"
// )

var (
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string

	API_PORT       string
	TOKEN_GRPC_URL string

	SHARED_LINK_ENDPOINT string
)

// LoadEnv Load Environment Variable
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		log.Println("using default environment variable")
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")

	API_PORT = os.Getenv("API_PORT")
	TOKEN_GRPC_URL = os.Getenv("TOKEN_GRPC_URL")

	SHARED_LINK_ENDPOINT = os.Getenv("SHARED_LINK_ENDPOINT")

}
