package rest

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/pebruwantoro/monorepo_project/backend/internal/app/container"
)

func StartRestHttpService(container *container.Container) {
	server := echo.New()

	SetupRouter(server, container)
	SetupMiddleware(server, container)

	// Start server
	go func() {
		if err := server.Start(fmt.Sprintf(":%d", container.Config.App.HttpPort)); err != nil {
			server.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}
