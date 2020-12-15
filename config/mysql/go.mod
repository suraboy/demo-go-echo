module github.com/suraboy/go-echo/config/mysql

go 1.15

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	github.com/suraboy/go-echo v1.0.8 // indirect
)

replace (
	github.com/suraboy/go-echo => ../../
	github.com/suraboy/go-echo/routes => ../../routes
)
