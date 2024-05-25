package rest

import (
	"github.com/labstack/echo/v4"

	"github.com/pebruwantoro/monorepo_project/backend/internal/app/container"
	"github.com/pebruwantoro/monorepo_project/backend/internal/app/handler/rest/health_check"
)

func SetupRouter(server *echo.Echo, container *container.Container) {
	// inject handler with usecase via container
	healthCheckHandler := health_check.NewHandler().Validate()

	server.GET("/health", healthCheckHandler.Check)

	v1 := server.Group("/v1")
	{
		v1.GET("/:id", healthCheckHandler.Check)
	}
}
