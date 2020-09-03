package godi_example

import "github.com/labstack/echo"

// IApi application api
type IApi interface {
	Start() error
}

// IController api controller
type IController interface {
	Name() string
	Register(group *echo.Group) error
	Middlewares() []echo.MiddlewareFunc
}
