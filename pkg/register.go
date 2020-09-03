package pkg

import (
	"github.com/noartem/godi"
	"github.com/noartem/godi-example/pkg/api"
	"github.com/noartem/godi-example/pkg/util"
)

// Register register factories
func Register(c *godi.Container) error {
	return c.RegisterCompose(
		api.Register,
		util.Register,
	)
}
