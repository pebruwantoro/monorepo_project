package health_check

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/monorepo_project/backend/internal/pkg/response"
)

func (h *handler) Check(c echo.Context) error {
	resp := response.DefaultResponse{
		Success: true,
		Message: "OK",
	}

	return c.JSON(http.StatusOK, resp)
}
