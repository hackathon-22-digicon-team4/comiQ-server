package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetAuthorsResponse struct {
	Authors []Author `json:"authors"`
}

type Author struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	ImageURL string `json:"imageUrl,omitempty"`
}

func (h *Handlers) GetAuthors(c echo.Context) error {
	ctx := c.Request().Context()
	authors, err := h.Repository.FindAllAuthors(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	as := make([]Author, 0)
	for _, author := range authors {
		as = append(as, Author{
			ID:       author.ID,
			Name:     author.Name,
			ImageURL: author.ImageURL(h.AssetHost),
		})
	}
	return c.JSON(http.StatusOK, GetAuthorsResponse{
		Authors: as,
	})
}
