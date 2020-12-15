package routes

import (
	"github.com/labstack/echo"
	"github.com/suraboy/go-echo/api"
)

func UserRoute(e *echo.Echo) {
	e.GET("/v1/users", api.GetAllUser)
	e.GET("/v1/users/:id", api.FindUser)
	e.POST("/v1/users", api.CreateUser)
	e.PUT("/v1/users/:id", api.UpdateUser)
	e.DELETE("/v1/users/:id", api.DeleteUser)
}
