package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func DbConfig() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db_name := os.Getenv("db_name")
	db_host := os.Getenv("db_host")
	db_user := os.Getenv("db_user")
	db_password := os.Getenv("db_password")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", db_user, db_password, db_host, db_name)
	return dsn
}
