package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.GET("/", hello)
	zipCodeRoute := e.GET("/:zipCode", getWeatherByZipCode)

	// Names
	zipCodeRoute.Name = "get-weather-by-zip-code"

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getWeatherByZipCode(c echo.Context) error {
	zipCode := c.Param("zipCode")

	weatherURL := fmt.Sprintf("api.openweathermap.org/data/2.5/weather?zip=%s&appid=%s", zipCode, API_KEY)
	// Request
	req, err := http.NewRequest("GET", weatherURL, nil)
	if err != nil {
		panic(err)
	}

	// Client
	client := &http.Client{}

	// Send HTTP request
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Defer the closing of the body
	defer res.Body.Close()

	fmt.Print(res)
	// Fill the data with the data from the JSON
	// var data Employee

	// Use json.Decode for reading streams of JSON data
	// if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
	// 	fmt.Println(err)
	// '}

	return c.JSON(http.StatusOK, "success")
}
