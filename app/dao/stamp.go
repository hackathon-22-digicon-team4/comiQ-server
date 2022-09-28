package dao

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/hackathon-22-digicon-team4/comiQ-server/gen/comiq/daocore"
)

func SelectAllStamps(ctx context.Context, tx *sql.Tx) ([]*daocore.Stamp, error) {
	sql, params, err := squirrel.Select(daocore.StampAllColumns...).From(daocore.StampTableName).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := tx.QueryContext(ctx, sql, params...)
	if err != nil {
		return nil, err
	}
	res := make([]*daocore.Stamp, 0)
	for rows.Next() {
		r := new(daocore.Stamp)
		if err := rows.Scan(r.Ptrs()...); err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func SelectStampsByIDs(ctx context.Context, tx *sql.Tx, IDs []string) ([]*daocore.Stamp, error) {
	query := squirrel.Select(daocore.StampAllColumns...).
		From(daocore.StampTableName).
		Where(squirrel.Eq{"id": IDs})
	sql, params, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := tx.QueryContext(ctx, sql, params...)
	if err != nil {
		return nil, err
	}
	res := make([]*daocore.Stamp, 0)
	for rows.Next() {
		r := new(daocore.Stamp)
		if err := rows.Scan(r.Ptrs()...); err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}
