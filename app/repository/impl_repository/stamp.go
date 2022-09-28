package impl_repository

import (
	"context"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/dao"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/dao/parser"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
)

func (r *Repository) FindAllStamps(ctx context.Context) ([]model.Stamp, error) {
	txn, err := r.db.BeginROTx(ctx)
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()
	stamps, err := dao.SelectAllStamps(ctx, txn)
	if err != nil {
		return nil, err
	}
	res := make([]model.Stamp, 0)
	for _, s := range stamps {
		res = append(res, parser.Stamp(s))
	}
	return res, nil
}

func (r *Repository) FindStampsByIDs(ctx context.Context, IDs []string) ([]model.Stamp, error) {
	txn, err := r.db.BeginROTx(ctx)
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()
	stamps, err := dao.SelectStampsByIDs(ctx, txn, IDs)
	if err != nil {
		return nil, err
	}
	res := make([]model.Stamp, 0)
	for _, s := range stamps {
		res = append(res, parser.Stamp(s))
	}
	return res, nil
}
