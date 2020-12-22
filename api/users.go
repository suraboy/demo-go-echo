package api

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	mysql "github.com/suraboy/go-echo/config"
	"github.com/suraboy/go-echo/models"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var messageError struct {
	Errors messageFormat `json:"errors"`
}

type messageFormat struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

//get user list
func GetAllUser(c echo.Context) (err error) {
	mysql := mysql.DbManager()
	var user []models.Users
	mysql.DB.Find(&user)
	return c.JSON(http.StatusOK, echo.Map{"data": user})
}

//find uset by id
func FindUser(c echo.Context) (err error) {
	mysql := mysql.DbManager()
	id := c.Param("id")
	user := models.Users{}

	if err := mysql.DB.Find(&user, id).Error; err != nil || mysql.DB.Find(&user, id).RowsAffected == 0 {
		var msgError messageFormat
		if mysql.DB.Find(&user, id).RowsAffected == 0 {
			msgError.StatusCode = http.StatusNotFound
			msgError.Message = "Not Found"
		} else {
			msgError.StatusCode = http.StatusInternalServerError
			msgError.Message = "Internal Server Error"
			msgError.Error = err.Error()
		}
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": user})
}

//create user
func CreateUser(c echo.Context) (err error) {
	mysql := mysql.DbManager()
	user := new(models.Users)
	var msgError messageFormat
	if err = c.Bind(user); err != nil {
		msgError.StatusCode = http.StatusBadRequest
		msgError.Message = "Bad Request"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}
	if err = c.Validate(user); err != nil {
		msgError.StatusCode = http.StatusUnprocessableEntity
		msgError.Message = "The given data was invalid."
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
	if err := mysql.DB.Create(&user).Error; err != nil {
		msgError.StatusCode = http.StatusExpectationFailed
		msgError.Message = "Expectation Failed"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	} // pass pointer of data to Create

	return c.JSON(http.StatusCreated, echo.Map{"data": user})
}

//update user
func UpdateUser(c echo.Context) (err error) {
	pass := ""
	mysql := mysql.DbManager()
	id := c.Param("id")
	user := models.Users{}
	var msgError messageFormat

	if err := c.Bind(&user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if user.Password != "" {
		password := []byte(user.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		pass = string(hashedPassword)
	}

	if err := mysql.DB.Find(&user, id).Error; err != nil {
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}

	if pass != "" {
		user.Password = pass
	}

	if err := mysql.DB.Save(&user).Error; err != nil {
		msgError.StatusCode = http.StatusExpectationFailed
		msgError.Message = "Expectation Failed"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(http.StatusExpectationFailed, messageError)
	}
	return c.JSON(http.StatusOK, echo.Map{"data": user})
}

//delete user
func DeleteUser(c echo.Context) (err error) {
	id := c.Param("id")
	mysql := mysql.DbManager()
	user := models.Users{}
	var msgError messageFormat
	if err := mysql.DB.Find(&user, id).Error; err != nil {
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}

	if err := mysql.DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.NoContent(http.StatusNoContent)
}

type RequestLogin struct {
	Username string `json:"username" sql:"username"`
	Password string `json:"password" sql:"username"`
}

type ResponseToken struct {
	TokenType string `json:"token_type"`
	ExpiresIn int64 `json:"expires_in"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func LoginUser(c echo.Context) (err error) {
	var msgError messageFormat
	var response ResponseToken
	mysql := mysql.DbManager()
	req := RequestLogin{}
	user := models.Users{}

	if err := c.Bind(&req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := mysql.DB.Where("username = ?", req.Username).Where("password = ?", req.Password).Find(&user).Error; err != nil {
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.JSON(msgError.StatusCode, messageError)
	}

	////compare the user from the request, with the one we defined:
	//if user.Username != user.Username || user.Password != user.Password {
	//	return c.JSON(http.StatusUnauthorized, "Please provide valid login details")
	//}
	token, err := CreateToken(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	response.AccessToken = token
	response.ExpiresIn = time.Now().Add(time.Minute * 15).Unix()
	response.TokenType = "Bearer"
	return c.JSON(http.StatusOK,response)
}

func CreateToken(user models.Users) (string, error) {
	var err error
	//key, err := loadKey("letsencrypt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = user.ID
	atClaims["name"] = user.Name
	atClaims["last_name"] = user.LastName
	atClaims["type"] = user.Type
	atClaims["created_at"] = user.CreatedAt
	atClaims["authorized"] = true
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

// loadKey reads and parses a private RSA key.
func loadKey(file string) (*rsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("pem decode: no key found")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
