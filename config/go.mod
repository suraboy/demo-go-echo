module github.com/suraboy/go-echo/config

go 1.15

replace (
	github.com/suraboy/go-echo => ../
	github.com/suraboy/go-echo/models => ../models
	github.com/suraboy/go-echo/routes => ../routes
)

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	github.com/suraboy/go-echo/models v0.0.0-00010101000000-000000000000
)
