package main

import (
	"fmt"
	"github.com/noartem/godi"
	types "github.com/noartem/godi-example"
	"github.com/noartem/godi-example/pkg"
)

/*
	[*] APP (HTTP)

	[*] HTTP (CONFIG, CONTROLLERS)

	[*] DB (CONFIG)

	[*] CONFIG

	[ ] CONTROLLERS (CONFIG)

		[*] AUTH (DB, CONFIG)

		[ ] USER (DB)

		[ ] PROJECTS (DB, MODULES)

		[ ] MODULES

			[ ] NOTES (DB, CONFIG)

			[ ] BOARD (DB, CONFIG)

			[ ] CHAT (DB, CONFIG)
*/

func exec() error {
	c, err := godi.NewContainerWithLogging()
	if err != nil {
		return err
	}

	err = c.RegisterCompose(pkg.Register)
	if err != nil {
		return err
	}

	apiI, err := c.Get("godi_example.IApi")
	if err != nil {
		return err
	}

	apiBean, ok := apiI.(types.IApi)
	if !ok {
		return fmt.Errorf("invalid apiI type: %T", apiI)
	}

	err = apiBean.Start()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := exec(); err != nil {
		panic(err)
	}
}
