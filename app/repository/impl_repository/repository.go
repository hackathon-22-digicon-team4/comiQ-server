package impl_repository

import (
	"context"
	"database/sql"

	"github.com/hackathon-22-digicon-team4/comiQ-server/pkg/db"
)

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) BeginROTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginROTx(ctx)
}

func (r *Repository) BeginRWTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginRWTx(ctx)
}
