module github.com/suraboy/go-echo/api

go 1.15

replace (
	github.com/suraboy/go-echo => ../
	github.com/suraboy/go-echo/models => ../models
)

require (
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/suraboy/go-echo/config v0.0.0-20201215203512-71cefd9f942b
	github.com/suraboy/go-echo/models v0.0.0-00010101000000-000000000000
)
