package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetBookByIDResponse struct {
	ID           string `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	BookSeriesID string `json:"bookSeriesId,omitempty"`
	TotalPages   int    `json:"totalPages,omitempty"`
	ImageURL     string `json:"imageUrl,omitempty"`
}

func (h *Handlers) GetBookByID(c echo.Context) error {
	ctx := c.Request().Context()
	bookID := c.Param("id")
	book, err := h.Repository.FindBookByID(ctx, bookID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, GetBookByIDResponse{
		ID:           book.ID,
		Title:        book.Title,
		BookSeriesID: book.BookSeriesID,
		TotalPages:   book.TotalPages,
		ImageURL:     book.ImageURL(h.AssetHost),
	})
}
