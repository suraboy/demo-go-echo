package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/suraboy/go-echo/routes"
	_ "golang.org/x/crypto/bcrypt"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello Go-echo")
	})

	routes.UserRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}

