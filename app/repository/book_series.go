package repository

import (
	"context"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
)

type BookSeriesRepository interface {
	FindAllBookSeries(ctx context.Context) ([]model.BookSeries, error)
}
