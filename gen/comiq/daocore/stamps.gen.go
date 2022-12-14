// Code generated by "script/dtogen". DO NOT EDIT.
package daocore

import (
  "context"
  "database/sql"
  "strings"
  "time"

  "github.com/Masterminds/squirrel"
)

const StampTableName = "stamps"

var StampAllColumns = []string{
  "id",
  
  "name",
  
  "created_at",
  
  "updated_at",
  
}

var StampColumnsWOMagics = []string{
  "id",
  
  "name",
  
  
  
}

var StampPrimaryKeyColumns = []string{
  "id",
  
  
  
  
}



type Stamp struct {
  ID string
  Name string
  CreatedAt *time.Time
  UpdatedAt *time.Time
}

func (t *Stamp) Values() []interface{} {
  return []interface{}{t.ID,t.Name,
  }
}

func (t *Stamp) SetMap() map[string]interface{} {
  return map[string]interface{}{"id": t.ID,"name": t.Name,
  }
}


func (t *Stamp) Ptrs() []interface{} {
  return []interface{}{
    &t.ID,
    &t.Name,
    &t.CreatedAt,
    &t.UpdatedAt,
  }
}




func IterateStamp(sc interface{ Scan(...interface{}) error}) (Stamp, error) {
  t := Stamp{}
  if err := sc.Scan(t.Ptrs()...); err != nil {
    return Stamp{}, MapError(err)
  }
  return t, nil
}

func SelectOneStampByID(ctx context.Context, txn *sql.Tx, id string) (Stamp, error) {
  query, params, err := squirrel.
    Select(StampAllColumns...).
    From(StampTableName).
    Where(squirrel.Eq{
      "id": id,
    }).
    ToSql()
  if err != nil {
    return Stamp{}, MapError(err)
  }
  stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
    return Stamp{}, MapError(err)
	}
  return IterateStamp(stmt.QueryRowContext(ctx, params...))
  
}




func InsertStamp(ctx context.Context, txn *sql.Tx, records []*Stamp) error {
  for i := range records {
		if records[i] == nil {
			records = append(records[:i], records[i+1:]...)
		}
	}
	if len(records) == 0 {
    return nil
  }
  sq := squirrel.Insert(StampTableName).Columns(StampColumnsWOMagics...)
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

func UpdateStamp(ctx context.Context, txn *sql.Tx, record Stamp) error {
	sql, params, err := squirrel.Update(StampTableName).SetMap(record.SetMap()).
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

func UpsertStamp(ctx context.Context, txn *sql.Tx, record Stamp) error {
	updateSQL, params, err := squirrel.Update(StampTableName).SetMap(record.SetMap()).ToSql()
	if err != nil {
		return err
	}
	updateSQL = strings.TrimPrefix(updateSQL, "UPDATE "+StampTableName+" SET ")
	query, params, err := squirrel.Insert(StampTableName).Columns(StampColumnsWOMagics...).Values(record.Values()...).SuffixExpr(squirrel.Expr("ON DUPLICATE KEY UPDATE "+updateSQL, params...)).ToSql()
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

func DeleteOneStampByID(ctx context.Context, txn *sql.Tx, id string) error {
  query, params, err := squirrel.
    Delete(StampTableName).
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

