module github.com/suraboy/go-echo/routes

go 1.15

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.1.17 // indirect
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/suraboy/go-echo v1.0.8 // indirect
	github.com/uniplaces/carbon v0.1.6 // indirect
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392
	gorm.io/driver/mysql v1.0.3 // indirect
	gorm.io/gorm v1.20.7
)

replace (
    github.com/suraboy/go-echo => ../
    github.com/suraboy/go-echo/models => ../models
    github.com/suraboy/go-echo/config => ../config
)
