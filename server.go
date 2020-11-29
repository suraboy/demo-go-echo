package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello Go-echo")
	})

	h := Handler{}
	h.connectDB()

	e.GET("/v1/users", h.GetAllUser)
	e.Logger.Fatal(e.Start(":8080"))
}

type (
	Handler struct {
		DB *gorm.DB
	}
)

//ให้เชื่อมต่อฐานข้อมูลเมื่อ Initialize
func (h *Handler) connectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := username + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	h.DB = db
}

// User
type Users struct {
	Id    uint    `json:"id"`
	Username    string   `json:"username"`
	Name  string `json:"name" xml:"name"`
	LastName string `json:"last_name"`
	Email string `json:"email" xml:"email"`
	Verify string `json:"verify"`
	Mobile string `json:"mobile"`
	Type	string `json:"type"`
	Pin string `json:"pin"`
	Status string `json:"status"`
	UserGroupId int `json:"user_group_id"`
	Gender string `json:"gender"`
	Birthday string `json:"birthday"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type messageError struct {
	Errors messageFormat `json:"errors"`
}

type messageFormat struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

var TransformUser struct {
	Data []Users `json:"data"`
}

func (h *Handler) GetAllUser(c echo.Context) error {
	//id := c.Param("id")
	rows, err := h.DB.Model(&Users{}).Rows()
	defer rows.Close()

	if err != nil {
		errorJson := `{"errors":{"status_code":404,"message":"Not Found."}}`
		var errors messageError
		json.Unmarshal([]byte(errorJson), &errors)

		return c.JSON(http.StatusNotFound, errors)
	}

	for rows.Next() {
		var userData Users
		// ScanRows scan a row into user
		err := h.DB.ScanRows(rows, &userData)
		if err != nil {
			panic(err.Error())
		}
		TransformUser.Data = append(TransformUser.Data,userData)
	}

	return c.JSON(200, TransformUser)
}
