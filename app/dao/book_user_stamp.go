package dao

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/hackathon-22-digicon-team4/comiQ-server/gen/comiq/daocore"
)

// BookID, BookSeriesID, UserID, StampIDでユーザーのスタンプを取得する
// ただし、それぞれは必須ではない
// 空文字の場合は、その条件は無視される
func SelectBookUserStampByQuery(ctx context.Context, tx *sql.Tx, bookSeriesID string, bookID string, userID string, stampID string) ([]*daocore.BookUserStamp, error) {
	query := squirrel.Select(daocore.BookUserStampAllColumns...).
		From(daocore.BookUserStampTableName)
	if bookSeriesID != "" {
		query = query.Where(squirrel.Eq{"book_series_id": bookSeriesID})
	}
	if bookID != "" {
		query = query.Where(squirrel.Eq{"book_id": bookID})
	}
	if userID != "" {
		query = query.Where(squirrel.Eq{"user_id": userID})
	}
	if stampID != "" {
		query = query.Where(squirrel.Eq{"stamp_id": stampID})
	}
	sql, params, err := query.ToSql()
	rows, err := tx.QueryContext(ctx, sql, params...)
	if err != nil {
		return nil, err
	}
	res := make([]*daocore.BookUserStamp, 0)
	for rows.Next() {
		r := new(daocore.BookUserStamp)
		if err := rows.Scan(r.Ptrs()...); err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}
