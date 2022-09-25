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
