package dao

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/hackathon-22-digicon-team4/comiQ-server/gen/comiq/daocore"
)

func SelectAllBookSeries(ctx context.Context, tx *sql.Tx) ([]*daocore.BookSery, error) {
	sql, params, err := squirrel.Select(daocore.BookSeryAllColumns...).From(daocore.BookSeryTableName).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := tx.QueryContext(ctx, sql, params...)
	if err != nil {
		return nil, err
	}
	res := make([]*daocore.BookSery, 0)
	for rows.Next() {
		r := new(daocore.BookSery)
		if err := rows.Scan(r.Ptrs()...); err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}
