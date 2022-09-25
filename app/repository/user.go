package repository

import (
	"context"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	FindUserByID(ctx context.Context, id string) (model.User, error)
}
