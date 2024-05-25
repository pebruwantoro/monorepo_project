package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DefaultResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccess(c echo.Context, data interface{}) error {
	resp := DefaultResponse{
		Success: true,
		Message: "Success",
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}
