package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookSeries struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"name,omitempty"`
	AuthorID    string `json:"authorId,omitempty"`
	AuthorName  string `json:"authorName,omitempty"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"imageUrl,omitempty"`
}

func (h *Handlers) GetBookSeries(c echo.Context) error {
	ctx := c.Request().Context()
	bookSeries, err := h.Repository.FindAllBookSeries(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	res := make([]BookSeries, 0)
	for _, b := range bookSeries {
		res = append(res, BookSeries{
			ID:          b.ID,
			Title:       b.Title,
			AuthorID:    b.AuthorID,
			AuthorName:  b.AuthorName,
			Description: b.Description,
			ImageURL:    b.ImageURL(h.AssetHost),
		})
	}
	return c.JSON(http.StatusOK, res)
}
