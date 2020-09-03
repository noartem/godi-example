package ent

import (
	"context"
	"fmt"
	"github.com/noartem/godi"
	"github.com/noartem/godi-example/ent"
	"github.com/noartem/godi-example/ent/migrate"
	"github.com/noartem/godi-example/pkg/util/config"

	_ "github.com/lib/pq"
)

type ClientWithCtx struct {
	DB *ent.Client
	Ctx context.Context
}

// NewEnt create new ent client
func NewEnt(config *config.Config) (*ClientWithCtx, *godi.BeanOptions, error) {
	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s",
			config.DB.Host, config.DB.Port, config.DB.User, config.DB.Name, config.DB.Password,
		),
	)
	if err != nil {
		return nil, nil, err
	}
	defer client.Close()

	ctx := context.Background()

	// Run the auto migration tool.
	err = client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true))
	if err != nil {
		return nil, nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	options := &godi.BeanOptions{
		Type: godi.Singleton,
	}

	clientWithCtx := &ClientWithCtx{
		DB:  client,
		Ctx: ctx,
	}

	return clientWithCtx, options, nil
}
