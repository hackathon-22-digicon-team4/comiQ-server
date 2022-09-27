package model

import (
	"fmt"
)

type Stamp struct {
	ID   string
	Name string
}

func (s Stamp) ImageURL(assetHost string) string {
	return fmt.Sprintf("https://%s/stamps/%s.jpeg", assetHost, s.ID)
}
