package routes

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	e := echo.New()
	h := Handler{}
	h.connectDB()
	e.GET("/v1/users", h.GetAllUser)
	e.GET("/v1/users/:id", h.FindUser)
	e.POST("/v1/users", h.CreateUser)
}

type (
	Handler struct {
		DB *gorm.DB
	}
)

//connect database mysql by gorm
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

	dsn := username + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&loc=Asia%2FBangkok&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	h.DB = db
}
// User
type Users struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Name        string    `json:"name" xml:"name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email" xml:"email"`
	Verify      string    `json:"verify" gorm:"type:enum('waiting','yes','no');default:'waiting'"`
	Mobile      string    `json:"mobile"`
	Type        string    `json:"type" gorm:"type:enum('owner','staff','other','admin','customer','brand-owner');default:'other'"`
	Pin         string    `json:"pin"`
	Status      string    `json:"status" gorm:"type:enum('active', 'inactive', 'ban');default:'inactive'"`
	UserGroupId int       `json:"user_group_id"`
	Gender      string    `json:"gender" gorm:"type:enum('male', 'female');default:'male'"`
	Birthday    time.Time `json:"birthday"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var messageError struct {
	Errors messageFormat `json:"errors"`
}

type messageFormat struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}
//get user list
func (h *Handler) GetAllUser(c echo.Context) (err error) {
	var user []Users
	h.DB.Find(&user)
	return c.JSON(http.StatusOK, echo.Map{"data": user})
}
//find uset by id
func (h *Handler) FindUser(c echo.Context) (err error) {
	id := c.Param("id")
	user := Users{}

	if err := h.DB.Find(&user, id).Error; err != nil || h.DB.Find(&user, id).RowsAffected == 0 {
		var msgError messageFormat
		if h.DB.Find(&user, id).RowsAffected == 0 {
			msgError.StatusCode = http.StatusNotFound
			msgError.Message = "Not Found"
		}else{
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
func (h *Handler) CreateUser(c echo.Context) (err error) {
	user := new(Users)
	if err = c.Bind(user); err != nil {
		var msgError messageFormat
		msgError.StatusCode = http.StatusBadRequest
		msgError.Message = "Bad Request"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(http.StatusBadRequest, messageError)
	}
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
	if err := h.DB.Create(&user).Error; err != nil {
		var msgError messageFormat
		msgError.StatusCode = http.StatusExpectationFailed
		msgError.Message = "Expectation Failed"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.JSON(http.StatusExpectationFailed, messageError)
	} // pass pointer of data to Create

	return c.JSON(http.StatusCreated, echo.Map{"data": user})
}