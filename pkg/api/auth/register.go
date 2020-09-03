package auth

import (
	"github.com/noartem/godi"
	"github.com/noartem/godi-example/pkg/api/auth/platform"
	"github.com/noartem/godi-example/pkg/api/auth/service"
	"github.com/noartem/godi-example/pkg/api/auth/transport"
)

// Register register service
func Register(c *godi.Container) error {
	return c.Register(
		service.NewService,
		platform.NewUserDB,
		transport.NewAuth,
	)
}