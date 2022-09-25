package model

import (
	"fmt"
)

type Book struct {
	ID           string
	Title        string
	BookSeriesID string
	TotalPages   int
}

func (b Book) ImageURL(assetHost string) string {
	return fmt.Sprintf("https://%s/books/%s.jpg", assetHost, b.ID)
}
