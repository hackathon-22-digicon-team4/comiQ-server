package model

import (
	"fmt"
)

type BookUserStamp struct {
	ID      string
	BookID  string
	BookSeriesID string
	PageNum int
	X       int
	Y       int
	UserID  string
	StampID string
}

func (b BookUserStamp) BookPageImageURL(assetHost string) string {
	return fmt.Sprintf("https://%s/books/%s/%d.jpg", assetHost, b.BookID, b.PageNum)
}
