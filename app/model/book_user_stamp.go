package model

import (
	"fmt"

	"github.com/google/uuid"
)

type BookUserStamp struct {
	ID           string
	BookID       string
	BookSeriesID string
	PageNum      int
	X            int
	Y            int
	UserID       string
	StampID      string
}

func (b BookUserStamp) BookPageImageURL(assetHost string) string {
	return fmt.Sprintf("https://%s/book_series/%s/books/%s/%d.jpeg", assetHost, b.BookSeriesID, b.BookID, b.PageNum)
}

func NewBookUserStamp(bookID, bookSeriesID string, pageNum, x, y int, userID, stampID string) BookUserStamp {
	return BookUserStamp{
		ID:           uuid.New().String(),
		BookID:       bookID,
		BookSeriesID: bookSeriesID,
		PageNum:      pageNum,
		X:            x,
		Y:            y,
		UserID:       userID,
		StampID:      stampID,
	}
}
