package api

import (
	"github.com/noartem/godi"
	"github.com/noartem/godi-example/pkg/api/auth"
)

func registerApi(c *godi.Container) error {
	return c.Register(
		NewApi,
	)
}

// Register register factories
func Register(c *godi.Container) error {
	return c.RegisterCompose(
		registerApi,

		auth.Register,
	)
}
