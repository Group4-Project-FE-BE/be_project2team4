package main

import (
	"be_project2team4/config"
	dComment "be_project2team4/feature/comment/delivery"
	rComment "be_project2team4/feature/comment/repository"
	sComment "be_project2team4/feature/comment/services"
	dPosting "be_project2team4/feature/posting/delivery"
	rPosting "be_project2team4/feature/posting/repository"
	sPosting "be_project2team4/feature/posting/services"
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

	mdlUser := rUser.New(db)
	mdlPosting := rPosting.New(db)
	mdlComment := rComment.New(db)

	serUser := sUser.New(mdlUser)
	serPosting := sPosting.New(mdlPosting)
	serComment := sComment.New(mdlComment)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	// //e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	// }))

	dUser.New(e, serUser)
	dPosting.New(e, serPosting)
	dComment.New(e, serComment)

	log.Fatal(e.Start(":8000"))

}
