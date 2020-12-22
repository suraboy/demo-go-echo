package routes

import (
	"errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo"
	"github.com/suraboy/go-echo/api"
	"gopkg.in/go-playground/validator.v9"
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
	e.GET("/v1/users", api.GetAllUser)
	e.GET("/v1/users/:id", api.FindUser)
	e.POST("/v1/users", api.CreateUser)
	e.PUT("/v1/users/:id", api.UpdateUser)
	e.DELETE("/v1/users/:id", api.DeleteUser)
	e.POST("/v1/login",api.LoginUser)
}
