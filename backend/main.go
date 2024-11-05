package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/nursultan-maratov/Diploma.git/internal/handler"
	"github.com/nursultan-maratov/Diploma.git/internal/middleware"
	"github.com/nursultan-maratov/Diploma.git/internal/postgres"
	"github.com/nursultan-maratov/Diploma.git/internal/service"
	"log"
)

func main() {

	db, err := postgres.ConnectDefaultDataBase()
	if err != nil {
		log.Fatalf("Kaput db is not working %v", err)
	}
	service := service.NewService(db)

	if err != nil {
		log.Fatalf("Kaput factory is not working %v", err)
	}

	middleware := middleware.NewMiddleware(service.GetRepository().GetUserRepo())

	newHandler := handler.NewHandler(service.GetUserManager())

	e := echo.New()
	e.POST("/create-user", newHandler.CreateUsers, middleware.SetUserToContext)
	e.Logger.Fatal(e.Start(":80"))
}
