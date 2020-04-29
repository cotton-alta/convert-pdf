package main

import (
	"md-to-pdf-api/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api", controllers.GetPdf())

	e.Start(":3001")
}

