package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetBookSeriesResponse struct {
	BookSeries []BookSeries `json:"bookSeries"`
}

type BookSeries struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"name,omitempty"`
	AuthorID    string `json:"authorId,omitempty"`
	AuthorName  string `json:"authorName,omitempty"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"imageUrl,omitempty"`
}

type Book struct {
	ID           string `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	BookSeriesID string `json:"bookSeriesId,omitempty"`
	TotalPages   int    `json:"totalPages,omitempty"`
	ImageURL     string `json:"imageUrl,omitempty"`
}

func (h *Handlers) GetBookSeries(c echo.Context) error {
	ctx := c.Request().Context()
	bookSeries, err := h.Repository.FindAllBookSeries(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	bs := make([]BookSeries, 0)
	for _, b := range bookSeries {
		bs = append(bs, BookSeries{
			ID:          b.ID,
			Title:       b.Title,
			AuthorID:    b.AuthorID,
			AuthorName:  b.AuthorName,
			Description: b.Description,
			ImageURL:    b.ImageURL(h.AssetHost),
		})
	}
	return c.JSON(http.StatusOK, GetBookSeriesResponse{
		BookSeries: bs,
	})
}

func (h *Handlers) GetBooksByBookSeriesID(c echo.Context) error {
	ctx := c.Request().Context()
	bookSeriesID := c.Param("id")
	books, err := h.Repository.FindBooksByBookSeriesID(ctx, bookSeriesID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	res := make([]Book, 0)
	for _, b := range books {
		res = append(res, Book{
			ID:           b.ID,
			Title:        b.Title,
			BookSeriesID: b.BookSeriesID,
			TotalPages:   b.TotalPages,
			ImageURL:     b.ImageURL(h.AssetHost),
		})
	}
	return c.JSON(http.StatusOK, res)
}
