package config

import (
	"os"

	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	DBPort    uint
	DBUser    string
	DBPwd     string
	DBHost    string
	DBName    string
	JWTSecret string
}

func NewConfig() *AppConfig {
	cfg := initConfig()
	if cfg == nil {
		log.Fatal("Cannot run configuration setup")
		return nil
	}

	return cfg
}

func initConfig() *AppConfig {
	var app AppConfig

	// err := godotenv.Load("config.env")
	// if err != nil {
	// 	log.Error("config error :", err.Error())
	// 	return nil
	// }

	app.DBUser = os.Getenv("DB_USER")
	app.DBPwd = os.Getenv("DB_PWD")
	app.DBHost = os.Getenv("DB_HOST")
	// port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	// if err != nil {
	// 	log.Error("config error :", err.Error())
	// 	return nil
	// }
	// app.DBPort = uint(port)
	// app.DBName = os.Getenv("DB_NAME")
	// app.JWTSecret = os.Getenv("JWT_SECRET")

	return &app
}
