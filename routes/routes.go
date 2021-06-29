package routes

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/suraboy/go-echo/api"
	"gopkg.in/go-playground/validator.v9"
	"os"
)

type CustomValidator struct {
	trans     ut.Translator
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)
	msg := ""
	for _, v := range errs.Translate(cv.trans) {
		if msg != "" {
			msg += ", "
		}
		msg += v
	}
	return errors.New(msg)
}
func UserRoute(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}
	r := e.Group("/v1/users")
	config := middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("ACCESS_SECRET")),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", api.GetAllUser)
	r.GET("/:id", api.FindUser)
	r.POST("", api.CreateUser)
	r.PUT("/:id", api.UpdateUser)
	r.DELETE("/:id", api.DeleteUser)

	e.POST("/v1/login",api.LoginUser)
}
