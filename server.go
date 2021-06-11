package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	moodRoute := e.GET("/mood/:mood", getMood)

	// Names
	zipCodeRoute.Name = "get-weather-by-zip-code"
	moodRoute.Name = "get-mood"

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Structs
type coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type temp struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type clouds struct {
	All int `json:"all"`
}

type sys struct {
	Systype int    `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type forecast struct {
	Coord      coord     `json:"coord"`
	Weather    []weather `json:"weather"`
	Base       string    `json:"base"`
	Main       temp      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       wind      `json:"wind"`
	Clouds     clouds    `json:"clouds"`
	Dt         int       `json:"dt"`
	Sys        sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getWeatherByZipCode(c echo.Context) error {
	zipCode := c.Param("zipCode")
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?zip=%s&appid=%s", zipCode, apiKey)

	// Request
	res, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// Fill the data with the body from the JSON
	var f forecast

	jsonErr := json.Unmarshal(body, &f)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Print(f)

	return c.JSON(http.StatusOK, f)
}

func getMood(c echo.Context) error {
	mood := c.Param("mood")
	resp := fmt.Sprintf("You are in a %s mood. Thanks for sharing!", mood)
	return c.String(http.StatusOK, resp)
}
