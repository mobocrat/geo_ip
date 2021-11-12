package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	go incrementer()
	e := echo.New()
	e.Use(ipLocater)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/stats", func(c echo.Context) error {
		return c.JSON(http.StatusOK, locationStats)
	})
	e.Logger.Fatal(e.Start("0.0.0.0:1323"))
}
