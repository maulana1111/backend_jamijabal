package main

import (
	"backend_jamijabal/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	config.ConnectDB()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hellow, ")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
