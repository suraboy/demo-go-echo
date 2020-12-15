module github.com/suraboy/go-echo/routes

go 1.15

require (
	github.com/suraboy/go-echo v1.0.8 // indirect
)

replace (
    github.com/suraboy/go-echo => ../
    github.com/suraboy/go-echo/models => ../models
    github.com/suraboy/go-echo/routes => ../routes
)
