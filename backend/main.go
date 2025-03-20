package main

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
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
	newMiddleware := middleware.NewMiddleware(db, "salt")

	factory := service.NewService(db)

	if err != nil {
		log.Fatalf("Kaput factory is not working %v", err)
	}

	newHandler := handler.NewHandler(factory.GetUserManager(), factory.GetProductManager(), "salt")

	e := echo.New()

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	userGroup := e.Group("/user")
	userGroup.POST("/create", newHandler.CreateUsers)

	e.POST("/auth", newHandler.Auth)

	e.POST("/buy-product", newHandler.BuyProduct, newMiddleware.Auth)
	e.GET("/list-order", newHandler.ListOrder, newMiddleware.Auth)
	e.GET("/list-product", newHandler.ListProduct)
	e.Logger.Fatal(e.Start(":80"))
}
