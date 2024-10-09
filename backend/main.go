package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/nursultan-maratov/Diploma.git/internal/handler"
	"github.com/nursultan-maratov/Diploma.git/internal/postgres"
	"log"
)

func main() {

	_, err := postgres.ConnectDefaultDataBase()
	if err != nil {
		log.Fatalf("Kaput db is not working %v", err)
	}
	serviceHandler := handler.NewHandler()

	e := echo.New()
	e.GET("/hello-world", serviceHandler.HelloWorld)
	e.Logger.Fatal(e.Start(":80"))
}
