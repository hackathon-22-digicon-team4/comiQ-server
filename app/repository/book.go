package repository

import (
	"context"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
)

type BookRepository interface {
	FindBooksByBookSeriesID(ctx context.Context, bookSeriesID string) ([]model.Book, error)
	FindBookByID(ctx context.Context, ID string) (model.Book, error)
}
