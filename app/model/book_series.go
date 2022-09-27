package model

import (
	"fmt"
)

type BookSeries struct {
	ID          string
	Title       string
	AuthorID    string
	AuthorName  string
	Description string
}

func (b BookSeries) ImageURL(assetHost string) string {
	return fmt.Sprintf("https://%s/book_series/%s/book_series_image.jpeg", assetHost, b.ID)
}
