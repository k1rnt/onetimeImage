package main

import (
	"github.com/k1rnt/onetimeImage/config"
	"github.com/k1rnt/onetimeImage/infrastructure"
	"github.com/k1rnt/onetimeImage/interface/handler"
	"github.com/k1rnt/onetimeImage/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	healthRepository := infrastructure.NewHealthRepository(config.NewDB())
	healthUsecase := usecase.NewHealthUsecase(healthRepository)
	healthHandler := handler.NewHealthHandler(healthUsecase)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	health_router(e, healthHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func health_router(e *echo.Echo, hH handler.HealthHandler) {
	e.POST("/health", hH.Post())
}
