package api

import (
	"github.com/labstack/echo"
	mysql "github.com/suraboy/go-echo/config"
	"github.com/suraboy/go-echo/models"
	"net/http"
)

//get user list
func GetAllUser(c echo.Context) (err error) {
	mysql := mysql.DbManager()
	var user []models.Users
	mysql.DB.Find(&user)
	return c.JSON(http.StatusOK, echo.Map{"data": user})
}