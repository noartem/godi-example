// Godi-example - example http server by github.com/noartem/godi
//
// API Docs for godi-example v1
//
// 	 Terms Of Service:  N/A
//     Schemes: http
//     Version: 2.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Noskov Artem <nowasmawesome@gmail.com> https://github.com/noartem
//     Host: localhost:8080
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer: []
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package api

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/noartem/godi"
	types "github.com/noartem/godi-example"
	"github.com/noartem/godi-example/pkg/util/config"
)

// Api application api
type Api struct {
	e           *echo.Echo
	controllers []types.IController
	config      *config.Config
}

// NewApi create new application api
func NewApi(controllers []types.IController, config *config.Config) (types.IApi, *godi.BeanOptions, error) {
	e := echo.New()

	validator, err := NewValidator()
	if err != nil {
		return nil, nil, err
	}
	e.Validator = validator

	api := &Api{
		e:           e,
		controllers: controllers,
		config:      config.Get(),
	}

	options := &godi.BeanOptions{Type: godi.Singleton}

	return api, options, nil
}

// Start start api server
func (api *Api) Start() error {
	for _, controller := range api.controllers {
		err := controller.Register(api.e)
		if err != nil {
			return fmt.Errorf("%v register: %v", controller, err)
		}
	}

	api.e.Static("/swagger", api.config.SwaggerPath)

	err := api.e.Start(fmt.Sprintf(":%d", api.config.Port))
	if err != nil {
		return err
	}

	return nil
}
