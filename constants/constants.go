package constants

import (
	"os"
)

var (
	MYSQL_PORT     string
	MYSQL_USER     string
	MYSQL_HOST     string
	MYSQL_DATABASE string
	MYSQL_PASSWORD string
	DRIVER         string
	APP_PORT       string
)

func init() {
	MYSQL_PORT = os.Getenv("MYSQL_PORT")
	MYSQL_USER = os.Getenv("MYSQL_USER")
	MYSQL_HOST = os.Getenv("MYSQL_HOST")
	MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE")
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	DRIVER = os.Getenv("DRIVER")
	APP_PORT = os.Getenv("APP_PORT")
}
