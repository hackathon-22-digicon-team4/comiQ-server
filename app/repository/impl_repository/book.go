package impl_repository

import (
	"context"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/dao"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/dao/parser"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
)

func (r *Repository) FindBooksByBookSeriesID(ctx context.Context, bookSeriesID string) ([]model.Book, error) {
	txn, err := r.db.BeginROTx(ctx)
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()
	books, err := dao.SelectBookByBookSeriesID(ctx, txn, bookSeriesID)
	if err != nil {
		return nil, err
	}
	res := make([]model.Book, 0)
	for _, b := range books {
		res = append(res, parser.Book(b))
	}
	return res, nil
}

func (r *Repository) FindBookByID(ctx context.Context, ID string) (model.Book, error) {
	txn, err := r.db.BeginROTx(ctx)
	if err != nil {
		return model.Book{}, err
	}
	defer txn.Rollback()
	book, err := dao.SelectOneBookByID(ctx, txn, ID)
	if err != nil {
		return model.Book{}, err
	}
	return parser.Book(&book), nil
}
