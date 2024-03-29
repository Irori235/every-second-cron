package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${uri} ${method} ${status} ${time_rfc3339}\n",
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
