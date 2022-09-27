package model

import (
	"fmt"
)

type Author struct {
	ID   string
	Name string
}

func (a Author) ImageURL(assetHost string) string {
	return fmt.Sprintf("https://%s/authors/%s.jpeg", assetHost, a.ID)
}
