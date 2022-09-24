package db

import (
	"context"
	"database/sql"
)

type DB struct {
	RoDB *sql.DB
	RwDB *sql.DB
}

func NewDB(roDB, rwDB *sql.DB) *DB {
	return &DB{
		RoDB: roDB,
		RwDB: rwDB,
	}
}

func (d *DB) BeginROTx(ctx context.Context) (*sql.Tx, error) {
	return d.RoDB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  true,
	})
}

func (d *DB) BeginRWTx(ctx context.Context) (*sql.Tx, error) {
	return d.RwDB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
}
