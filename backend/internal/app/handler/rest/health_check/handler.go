package health_check

import "github.com/labstack/echo/v4"

type HealthCheckHandler interface {
	Check(c echo.Context) error
}

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Validate() HealthCheckHandler {
	return h
}
