package config

import (
	"os"
	"strconv"

	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	DBPort   uint
	DBUser   string
	DBPwd    string
	DBHost   string
	DBName   string
	JWSecret string
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
	// isProduction, errPro := strconv.Atoi(os.Getenv("IS_PRODuCTION"))
	// if errPro != nil {
	// 	log.Error("config error :", errPro.Error())
	// 	return nil
	// }

	// err := godotenv.Load("config.env")
	// if err != nil {
	// 	log.Error("config error :", err.Error())
	// 	return nil
	// }

	app.DBUser = os.Getenv("DB_USER")
	app.DBPwd = os.Getenv("DB_PWD")
	app.DBHost = os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Error("config error :", err.Error())
		return nil
	}
	app.DBPort = uint(port)
	app.DBName = os.Getenv("DB_NAME")
	app.JWSecret = os.Getenv("JW_SECRET")

	return &app
}
