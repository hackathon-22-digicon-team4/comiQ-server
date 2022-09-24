package repository

import (
	"context"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) error
	FindByID(ctx context.Context, id string) (model.User, error)
}
