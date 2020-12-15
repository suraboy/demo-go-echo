module github.com/suraboy/go-echo/routes

go 1.15

replace (
	github.com/suraboy/go-echo => ../
	github.com/suraboy/go-echo/config => ../config
	github.com/suraboy/go-echo/models => ../models
	github.com/suraboy/go-echo/api => ../api
)

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/suraboy/go-echo/config v0.0.0-00010101000000-000000000000
	github.com/suraboy/go-echo/models v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9
)
