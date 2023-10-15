package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var config *Config

type Config struct {
	ProductSvcHost         string
	ListProductUri         string
	CreateProductUri       string
	GetProductUri          string
	UpdateProductUri       string
	DisableProductUri      string
	EnableProductUri       string
	IncreaseBookedQuotaUri string
	DecreaseBookedQuotaUri string

	AuthSvcHost     string
	LoginUri        string
	LogoutUri       string
	AuthenticateUri string

	OrderSvcHost   string
	AddToCartUri   string
	CreateOrderUri string
	CancelOrderUri string
	ListOrdersUri  string
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config = &Config{
		ProductSvcHost:         os.Getenv("PRODUCT_SERVICE_HOST"),
		ListProductUri:         os.Getenv("LIST_PRODUCT_URI"),
		CreateProductUri:       os.Getenv("CREATE_PRODUCT_URI"),
		GetProductUri:          os.Getenv("GET_PRODUCT_URI"),
		UpdateProductUri:       os.Getenv("UPDATE_PRODUCT_URI"),
		DisableProductUri:      os.Getenv("DISABLE_PRODUCT_URI"),
		EnableProductUri:       os.Getenv("ENABLE_PRODUCT_URI"),
		IncreaseBookedQuotaUri: os.Getenv("INCREASE_BOOKED_QUOTA"),
		DecreaseBookedQuotaUri: os.Getenv("DECREASE_BOOKED_QUOTA"),

		AuthSvcHost:     os.Getenv("AUTH_SERVICE_HOST"),
		LoginUri:        os.Getenv("LOGIN_URI"),
		LogoutUri:       os.Getenv("LOGOUT_URI"),
		AuthenticateUri: os.Getenv("AUTHENTICATE_URI"),

		OrderSvcHost:   os.Getenv("ORDER_SERVICE_HOST"),
		AddToCartUri:   os.Getenv("ADD_TO_CART_URI"),
		CreateOrderUri: os.Getenv("CREATE_ORDER_URI"),
		CancelOrderUri: os.Getenv("CANCEL_ORDER_URI"),
		ListOrdersUri:  os.Getenv("LIST_ORDERS_URI"),
	}

	return config
}

func GetConfig() *Config {
	return config
}
