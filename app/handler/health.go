package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ヘルスチェック
func (h *Handlers) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
