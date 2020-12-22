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
	github.com/ericchiang/letsencrypt v0.0.0-20180212195347-0367c87bed38 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/labstack/echo v3.3.10+incompatible
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/square/go-jose v2.5.1+incompatible // indirect
	github.com/suraboy/go-echo/api v0.0.0-00010101000000-000000000000 // indirect
	github.com/suraboy/go-echo/models v0.0.0-00010101000000-000000000000
	github.com/suraboy/go-echo/routes v0.0.0-00010101000000-000000000000
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
)
