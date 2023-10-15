package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var config *Config

type Config struct {
	ProductSvcHost string

	AuthSvcHost     string
	LoginUri        string
	LogoutUri       string
	AuthenticateUri string

	OrderSvcHost string
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config = &Config{
		ProductSvcHost: os.Getenv("PRODUCT_SERVICE_HOST"),

		AuthSvcHost:     os.Getenv("AUTH_SERVICE_HOST"),
		LoginUri:        os.Getenv("LOGIN_URI"),
		LogoutUri:       os.Getenv("LOGOUT_URI"),
		AuthenticateUri: os.Getenv("AUTHENTICATE_URI"),

		OrderSvcHost: os.Getenv("ORDER_SERVICE_HOST"),
	}

	return config
}

func GetConfig() *Config {
	return config
}
