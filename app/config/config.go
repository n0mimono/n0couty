package config

import "os"

var PORT string

var DB_NAME string
var DB_USER string
var DB_PASSWORD string
var DB_HOST string
var DB_PORT string

var ML_HOST string
var ML_PORT string

func Init() {
	PORT = os.Getenv("PORT")

	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")

	ML_HOST = os.Getenv("ML_HOST")
	ML_PORT = os.Getenv("ML_PORT")
}
