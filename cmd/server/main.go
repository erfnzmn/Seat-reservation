package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Sample route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Seat Reservation API is running!")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}