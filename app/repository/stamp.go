package repository

import (
	"context"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
)

type StampRepository interface {
	FindAllStamps(ctx context.Context) ([]model.Stamp, error)
	FindStampsByIDs(ctx context.Context, IDs []string) ([]model.Stamp, error)
}
