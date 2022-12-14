// Code generated by "script/dtogen". DO NOT EDIT.
package daocore

import (
  "context"
  "database/sql"
  "strings"
  "time"

  "github.com/Masterminds/squirrel"
)

const BookTableName = "books"

var BookAllColumns = []string{
  "id",
  
  "title",
  
  "book_series_id",
  
  "total_pages",
  
  "created_at",
  
  "updated_at",
  
}

var BookColumnsWOMagics = []string{
  "id",
  
  "title",
  
  "book_series_id",
  
  "total_pages",
  
  
  
}

var BookPrimaryKeyColumns = []string{
  "id",
  
  
  
  
  
  
}



type Book struct {
  ID string
  Title string
  BookSeriesID string
  TotalPage int
  CreatedAt *time.Time
  UpdatedAt *time.Time
}

func (t *Book) Values() []interface{} {
  return []interface{}{t.ID,t.Title,t.BookSeriesID,t.TotalPage,
  }
}

func (t *Book) SetMap() map[string]interface{} {
  return map[string]interface{}{"id": t.ID,"title": t.Title,"book_series_id": t.BookSeriesID,"total_pages": t.TotalPage,
  }
}


func (t *Book) Ptrs() []interface{} {
  return []interface{}{
    &t.ID,
    &t.Title,
    &t.BookSeriesID,
    &t.TotalPage,
    &t.CreatedAt,
    &t.UpdatedAt,
  }
}




func IterateBook(sc interface{ Scan(...interface{}) error}) (Book, error) {
  t := Book{}
  if err := sc.Scan(t.Ptrs()...); err != nil {
    return Book{}, MapError(err)
  }
  return t, nil
}

func SelectBookByBookSeriesID(ctx context.Context, txn *sql.Tx, book_series_id string) ([]*Book, error) {
  query, params, err := squirrel.
    Select(BookAllColumns...).
    From(BookTableName).
    Where(squirrel.Eq{
      "book_series_id": book_series_id,
    }).
    ToSql()
  if err != nil {
    return nil, MapError(err)
  }
  stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
    return nil, MapError(err)
	}
  rows, err := stmt.QueryContext(ctx, params...)
  if err != nil {
    return nil, MapError(err)
  }
  res := make([]*Book, 0)
  for rows.Next() {
    t, err := IterateBook(rows)
    if err != nil {
      return nil, MapError(err)
    }
    res = append(res, &t)
  }
  return res, nil
  
}

func SelectOneBookByID(ctx context.Context, txn *sql.Tx, id string) (Book, error) {
  query, params, err := squirrel.
    Select(BookAllColumns...).
    From(BookTableName).
    Where(squirrel.Eq{
      "id": id,
    }).
    ToSql()
  if err != nil {
    return Book{}, MapError(err)
  }
  stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
    return Book{}, MapError(err)
	}
  return IterateBook(stmt.QueryRowContext(ctx, params...))
  
}




func InsertBook(ctx context.Context, txn *sql.Tx, records []*Book) error {
  for i := range records {
		if records[i] == nil {
			records = append(records[:i], records[i+1:]...)
		}
	}
	if len(records) == 0 {
    return nil
  }
  sq := squirrel.Insert(BookTableName).Columns(BookColumnsWOMagics...)
	for _, r := range records {
		if r == nil {
			continue
		}
		sq = sq.Values(r.Values()...)
	}
	query, params, err := sq.ToSql()
	if err != nil {
		return err
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return MapError(err)
	}
	if _, err = stmt.Exec(params...); err != nil {
		return MapError(err)
	}
	return nil
}

func UpdateBook(ctx context.Context, txn *sql.Tx, record Book) error {
	sql, params, err := squirrel.Update(BookTableName).SetMap(record.SetMap()).
		Where(squirrel.Eq{
      "id": record.ID,
    }).
		ToSql()
	if err != nil {
		return err
	}
	stmt, err := txn.PrepareContext(ctx, sql)
	if err != nil {
		return MapError(err)
	}
	if _, err = stmt.Exec(params...); err != nil {
		return MapError(err)
	}
	return nil
}

func UpsertBook(ctx context.Context, txn *sql.Tx, record Book) error {
	updateSQL, params, err := squirrel.Update(BookTableName).SetMap(record.SetMap()).ToSql()
	if err != nil {
		return err
	}
	updateSQL = strings.TrimPrefix(updateSQL, "UPDATE "+BookTableName+" SET ")
	query, params, err := squirrel.Insert(BookTableName).Columns(BookColumnsWOMagics...).Values(record.Values()...).SuffixExpr(squirrel.Expr("ON DUPLICATE KEY UPDATE "+updateSQL, params...)).ToSql()
	if err != nil {
		return err
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return MapError(err)
	}
	if _, err = stmt.Exec(params...); err != nil {
		return MapError(err)
	}
	return nil
}

func DeleteBookByBookSeriesID(ctx context.Context, txn *sql.Tx, book_series_id string) error {
  query, params, err := squirrel.
    Delete(BookTableName).
    Where(squirrel.Eq{
      "book_series_id": book_series_id,
    }).
    ToSql()
  if err != nil {
    return MapError(err)
  }
  stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return MapError(err)
	}
	if _, err = stmt.Exec(params...); err != nil {
		return MapError(err)
	}
	return nil
}

func DeleteOneBookByID(ctx context.Context, txn *sql.Tx, id string) error {
  query, params, err := squirrel.
    Delete(BookTableName).
    Where(squirrel.Eq{
      "id": id,
    }).
    ToSql()
  if err != nil {
    return MapError(err)
  }
  stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return MapError(err)
	}
	if _, err = stmt.Exec(params...); err != nil {
		return MapError(err)
	}
	return nil
}

