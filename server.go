package main

import (
	"github.com/labstack/echo"
	//"github.com/suraboy/go-echo/api/routes"
)

func main() {
	e := echo.New()
	//routes.init()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World WTF!")
	//})
	e.Logger.Fatal(e.Start(":8080"))
}