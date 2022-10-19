package main

import (
	"be_project2team4/config"
	dUser "be_project2team4/feature/user/delivery"
	rUser "be_project2team4/feature/user/repository"
	sUser "be_project2team4/feature/user/services"
	"be_project2team4/utils/database"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//pemanggilan config
	cfg := config.NewConfig()
	db := database.InitDB(cfg)

	mdl := rUser.New(db)

	serUser := sUser.New(mdl)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	// //e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	// }))

	dUser.New(e, serUser)

	log.Fatal(e.Start(":8000"))

}
