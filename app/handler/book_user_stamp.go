package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
	"github.com/hackathon-22-digicon-team4/comiQ-server/pkg/echoutil"
)

type GetBookUserStampsResponse struct {
	BookUserStamps []BookUserStamp `json:"bookUserStamps"`
}

type BookUserStamp struct {
	ID               string `json:"id,omitempty"`
	BookID           string `json:"bookId,omitempty"`
	BookSeriesID     string `json:"bookSeriesId,omitempty"`
	PageNum          int    `json:"pageNum,omitempty"`
	X                int    `json:"x,omitempty"`
	Y                int    `json:"y,omitempty"`
	UserID           string `json:"userId,omitempty"`
	StampID          string `json:"stampId,omitempty"`
	BookPageImageURL string `json:"bookPageImageUrl,omitempty"`
}

type PostBookUserStampsRequest struct {
	BookID       string `json:"bookId,omitempty"`
	BookSeriesID string `json:"bookSeriesId,omitempty"`
	PageNum      int    `json:"pageNum,omitempty"`
	X            int    `json:"x,omitempty"`
	Y            int    `json:"y,omitempty"`
	StampID      string `json:"stampId,omitempty"`
}

type PostBookUserStampsResponse struct {
	ID 			 string `json:"id,omitempty"`
}

func (h *Handlers) GetBookUserStamps(c echo.Context) error {
	ctx := c.Request().Context()
	bookSeriesID := c.QueryParam("bookSeriesId")
	bookID := c.QueryParam("bookId")
	userID := c.QueryParam("userId")
	stampID := c.QueryParam("stampId")
	fmt.Println("bookSeriesID: ", bookSeriesID)
	bookUserStamps, err := h.Repository.FindBookUserStampsByQuery(ctx, bookSeriesID, bookID, userID, stampID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	buses := make([]BookUserStamp, 0)
	for _, bus := range bookUserStamps {
		buses = append(buses, BookUserStamp{
			ID:               bus.ID,
			BookID:           bus.BookID,
			BookSeriesID:     bus.BookSeriesID,
			PageNum:          bus.PageNum,
			X:                bus.X,
			Y:                bus.Y,
			UserID:           bus.UserID,
			StampID:          bus.StampID,
			BookPageImageURL: bus.BookPageImageURL(h.AssetHost),
		})
	}
	fmt.Println("buses: ", buses)
	return c.JSON(http.StatusOK, GetBookUserStampsResponse{
		BookUserStamps: buses,
	})
}

// Create a new book user stamp
func (h *Handlers) PostBookUserStamps(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(PostBookUserStampsRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	// userIDはsessionから取得する
	userID, err := echoutil.GetUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	bus := model.NewBookUserStamp(req.BookID, req.BookSeriesID, req.PageNum, req.X, req.Y, userID, req.StampID)
	if err := h.Repository.CreateBookUserStamp(ctx, bus); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, PostBookUserStampsResponse{
		ID: bus.ID,
	})
}
