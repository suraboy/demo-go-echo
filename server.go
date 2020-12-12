package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "golang.org/x/crypto/bcrypt"
	"net/http"
	_ "github.com/suraboy/go-echo/routes"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello Go-echo")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

