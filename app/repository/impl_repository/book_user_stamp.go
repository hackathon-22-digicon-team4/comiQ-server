package impl_repository

import (
	"context"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/dao"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/dao/parser"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
)

func (r *Repository) FindBookUserStampsByQuery(ctx context.Context, bookSeriesID string, bookID string, userID string, stampID string) ([]model.BookUserStamp, error) {
	txn, err := r.db.BeginROTx(ctx)
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()
	bookUserStamps, err := dao.SelectBookUserStampByQuery(ctx, txn, bookSeriesID, bookID, userID, stampID)
	if err != nil {
		return nil, err
	}
	res := make([]model.BookUserStamp, 0)
	for _, s := range bookUserStamps {
		res = append(res, parser.BookUserStamp(s))
	}
	return res, nil
}

func (r *Repository) CreateBookUserStamp(ctx context.Context, bookUserStamp model.BookUserStamp) error {
	txn, err := r.BeginRWTx(ctx)
	if err != nil {
		return err
	}
	defer txn.Rollback()
	err = dao.InsertBookUserStamp(ctx, txn, parser.BookUserStampModel{
		BookUserStamp: bookUserStamp,
		}.ToDao()); if err != nil {
		return err
	}
	return txn.Commit()
}

func (r *Repository) DeleteBookUserStampByID(ctx context.Context, id string) error {
	txn, err := r.BeginRWTx(ctx)
	if err != nil {
		return err
	}
	defer txn.Rollback()
	err = dao.DeleteOneBookUserStampByID(ctx, txn, id)
	if err != nil {
		return err
	}
	return txn.Commit()
}
