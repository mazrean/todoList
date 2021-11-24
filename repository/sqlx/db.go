package sqlx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	dbCtxKey = "db"
)

type DB struct {
	db *sqlx.DB
}

func NewDB() (*DB, error) {
	user, ok := os.LookupEnv("DB_USERNAME")
	if !ok {
		return nil, errors.New("DB_USERNAME is not set")
	}

	pass, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return nil, errors.New("DB_PASSWORD is not set")
	}

	host, ok := os.LookupEnv("DB_HOSTNAME")
	if !ok {
		return nil, errors.New("DB_HOSTNAME is not set")
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, errors.New("DB_PORT is not set")
	}

	database, ok := os.LookupEnv("DB_DATABASE")
	if !ok {
		return nil, errors.New("DB_DATABASE is not set")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FTokyo&charset=utf8mb4",
		user,
		pass,
		host,
		port,
		database,
	)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &DB{
		db: db,
	}, nil
}

func (db *DB) Transaction(ctx context.Context, txOption *sql.TxOptions, fn func(ctx context.Context) error) error {
	tx, err := db.db.BeginTxx(ctx, txOption)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	ctx = context.WithValue(ctx, dbCtxKey, tx)

	err = fn(ctx)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return fmt.Errorf("failed to rollback transaction: %w", err2)
		}

		return fmt.Errorf("failed in transaction: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

type db interface {
	sqlx.QueryerContext
	sqlx.ExecerContext
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func (db *DB) getDB(ctx context.Context) (db, error) {
	iDB := ctx.Value(dbCtxKey)
	if iDB == nil {
		return db.db, nil
	}

	sqlxDB, ok := iDB.(*sqlx.Tx)
	if !ok {
		return nil, errors.New("failed to get sqlx.Tx from context")
	}

	return sqlxDB, nil
}
