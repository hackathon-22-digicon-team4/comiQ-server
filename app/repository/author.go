package repository

import (
	"context"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
)

type AuthorRepository interface {
	FindAllAuthors(ctx context.Context) ([]model.Author, error)
}
