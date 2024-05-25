package container

import (
	"github.com/pebruwantoro/monorepo_project/backend/config"
	"github.com/pebruwantoro/monorepo_project/backend/internal/app/driver"
)

type Container struct {
	Config config.Config
}

func Setup() *Container {
	// Load Config
	cfg := config.Load()

	// Setup Driver
	_, _ = driver.NewPostgreSQLDatabase(cfg.DB)

	// Setup Repository

	// Setup Usecase

	return &Container{
		Config: cfg,
	}
}
