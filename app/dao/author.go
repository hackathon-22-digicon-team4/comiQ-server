package dao

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/hackathon-22-digicon-team4/comiQ-server/gen/comiq/daocore"
)

func SelectAllAuthors(ctx context.Context, tx *sql.Tx) ([]*daocore.Author, error) {
	sql, params, err := squirrel.Select(daocore.AuthorAllColumns...).From(daocore.AuthorTableName).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := tx.QueryContext(ctx, sql, params...)
	if err != nil {
		return nil, err
	}
	res := make([]*daocore.Author, 0)
	for rows.Next() {
		r := new(daocore.Author)
		if err := rows.Scan(r.Ptrs()...); err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}
