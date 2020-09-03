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
func NewApi(controllers []types.IController, config *config.Config) (types.IApi, *godi.BeanOptions) {
	e := echo.New()

	api := &Api{
		e:           e,
		controllers: controllers,
		config:      config.Get(),
	}

	options := &godi.BeanOptions{ Type: godi.Singleton }

	return api, options
}

// Start start api server
func (api *Api) Start() error {
	for _, controller := range api.controllers {
		group := api.e.Group(controller.Name(), controller.Middlewares()...)
		err := controller.Register(group)
		if err != nil {
			return fmt.Errorf("%v register: %v", group, err)
		}
	}

	err := api.e.Start(fmt.Sprintf(":%d", api.config.Port))
	if err != nil {
		return err
	}

	return nil
}
