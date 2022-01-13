package main

import (
	"Tribal/app/controllers"
	"Tribal/app/helpers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	helpers.PublicConfig = helpers.LoadConfig()
	helpers.DB = helpers.Connect(helpers.PublicConfig)

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Tribal! <3")
	})

	e.POST("/credit/validate", func(c echo.Context) error {
		creditController := controllers.NewCreditController()
		return creditController.Validate(c)
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
