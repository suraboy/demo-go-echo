module github.com/suraboy/go-echo

go 1.15

replace (
	github.com/suraboy/go-echo/api => ./api
	github.com/suraboy/go-echo/config => ./config
	github.com/suraboy/go-echo/models => ./models
	github.com/suraboy/go-echo/routes => ./routes
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/suraboy/go-echo/api v0.0.0-00010101000000-000000000000 // indirect
	github.com/suraboy/go-echo/routes v0.0.0-00010101000000-000000000000
)
