package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/nursultan-maratov/Diploma.git/internal/handler"
	"github.com/nursultan-maratov/Diploma.git/internal/postgres"
	"github.com/nursultan-maratov/Diploma.git/internal/service"
	"log"
)

func main() {

	db, err := postgres.ConnectDefaultDataBase()
	if err != nil {
		log.Fatalf("Kaput db is not working %v", err)
	}
	factory := service.NewService(db)

	if err != nil {
		log.Fatalf("Kaput factory is not working %v", err)
	}

	newHandler := handler.NewHandler(factory.GetUserManager())

	e := echo.New()
	userGroup := e.Group("/user")
	userGroup.POST("/create", newHandler.CreateUsers)
	e.Logger.Fatal(e.Start(":80"))
}
