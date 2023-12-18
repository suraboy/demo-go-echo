module github.com/suraboy/go-echo/routes

go 1.15

replace (
	github.com/suraboy/go-echo => ../
	github.com/suraboy/go-echo/api => ../api
	github.com/suraboy/go-echo/config => ../config
	github.com/suraboy/go-echo/models => ../models
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-playground/universal-translator v0.18.1
	github.com/labstack/echo v3.3.10+incompatible
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/suraboy/go-echo/api v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.17.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
)
