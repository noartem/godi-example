package util

import (
	"github.com/noartem/godi"
	"github.com/noartem/godi-example/pkg/util/config"
	"github.com/noartem/godi-example/pkg/util/ent"
	"github.com/noartem/godi-example/pkg/util/jwt"
)

func Register(c *godi.Container) error {
	return c.Register(
		config.Path,
		config.NewConfig,

		jwt.NewJWT,

		NewHash,

		ent.NewEnt,
	)
}
