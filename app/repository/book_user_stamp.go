package repository

import (
	"context"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
)

type BookUserStampRepository interface {
	FindBookUserStampsByQuery(ctx context.Context, bookSeriesID string, bookID string, users string, userID string, stampID string) ([]model.BookUserStamp, error)
	CreateBookUserStamp(ctx context.Context, bookUserStamp model.BookUserStamp) error
	DeleteBookUserStampByID(ctx context.Context, id string) error
}
