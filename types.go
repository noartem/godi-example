package godi_example

import "github.com/labstack/echo"

// IApi application api
type IApi interface {
	Start() error
}

// IController api controller
type IController interface {
	Register(e *echo.Echo) error
	Middlewares() []echo.MiddlewareFunc
}
