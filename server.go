package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/starlightromero/weather/handlers"
)

func main() {
	// Echo instance
	e := echo.New()

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
	}))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handlers.Hello)
	zipCodeRoute := e.GET("/:zipCode", handlers.GetWeatherByZipCode)
	moodRoute := e.GET("/mood/:mood", handlers.GetMood)

	// Names
	zipCodeRoute.Name = "get-weather-by-zip-code"
	moodRoute.Name = "get-mood"

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
