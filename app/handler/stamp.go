package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetStampsResponse struct {
	Stamps []Stamp `json:"stamps"`
}

type Stamp struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	ImageURL string `json:"imageUrl,omitempty"`
}

func (h *Handlers) GetStamps(c echo.Context) error {
	ctx := c.Request().Context()
	stamps, err := h.Repository.FindAllStamps(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	ss := make([]Stamp, 0)
	for _, stamp := range stamps {
		ss = append(ss, Stamp{
			ID:       stamp.ID,
			Name:     stamp.Name,
			ImageURL: stamp.ImageURL(h.AssetHost),
		})
	}
	return c.JSON(http.StatusOK, GetStampsResponse{
		Stamps: ss,
	})
}
