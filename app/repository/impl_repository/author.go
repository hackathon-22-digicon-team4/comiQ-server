package impl_repository

import (
	"context"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/dao"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/dao/parser"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
)

func (r *Repository) FindAllAuthors(ctx context.Context) ([]model.Author, error) {
	txn, err := r.db.BeginROTx(ctx)
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()
	authors, err := dao.SelectAllAuthors(ctx, txn)
	if err != nil {
		return nil, err
	}
	res := make([]model.Author, 0)
	for _, author := range authors {
		res = append(res, parser.Author(author))
	}
	return res, nil
}
