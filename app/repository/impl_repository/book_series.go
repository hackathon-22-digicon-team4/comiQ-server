package impl_repository

import (
	"context"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/dao"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/dao/parser"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
)

func (r *Repository) FindAllBookSeries(ctx context.Context) ([]model.BookSeries, error) {
	txn, err := r.db.BeginROTx(ctx)
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()
	bookSeries, err := dao.SelectAllBookSeries(ctx, txn)
	if err != nil {
		return nil, err
	}
	authors, err := dao.SelectAllAuthors(ctx, txn)
	if err != nil {
		return nil, err
	}
	authorID2Name := make(map[string]string)
	for _, author := range authors {
		authorID2Name[author.ID] = author.Name
	}
	res := make([]model.BookSeries, 0)
	for _, b := range bookSeries {
		res = append(res, parser.BookSeries(b, authorID2Name[b.AuthorID]))
	}
	return res, nil
}
