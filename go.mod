module github.com/suraboy/go-echo

go 1.15

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.1.17 // indirect
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/suraboy/go-echo/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/suraboy/go-echo/routes v0.0.0-00010101000000-000000000000
	github.com/uniplaces/carbon v0.1.6 // indirect
	golang.org/x/crypto v0.0.0-20201124201722-c8d3bf9c5392 // indirect
	golang.org/x/net v0.0.0-20201209123823-ac852fbbde11 // indirect
	golang.org/x/sys v0.0.0-20201211090839-8ad439b19e0f // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/genproto v0.0.0-20201211151036-40ec1c210f7a // indirect
	google.golang.org/grpc v1.34.0 // indirect
	gorm.io/driver/mysql v1.0.3 // indirect
	gorm.io/gorm v1.20.7 // indirect
)

replace (
	github.com/suraboy/go-echo/config => ./config
	github.com/suraboy/go-echo/routes => ./routes
)
