package database

import (
	"context"
	"database/sql"
	"os"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

// DBTransaction used to aggregate transactions
type DBTransaction struct {
	postgres *sql.Tx
	Builder  squirrel.StatementBuilderType
	ctx      context.Context
}

// Config used to configuration of database
type Config struct {
	Drive string
	Host  string
	Port  string
	User  string
	Pass  string
	Name  string
}

// OpenConnection initialize connection with database
func OpenConnection(ctx context.Context, readOnly bool) (*DBTransaction, error) {
	var (
		t   = &DBTransaction{}
		db  *sql.DB
		err error
		c   = Config{
			Drive: os.Getenv("DB_DRIVE"),
			Host:  os.Getenv("DB_HOST"),
			Port:  os.Getenv("DB_PORT"),
			User:  os.Getenv("DB_USER"),
			Pass:  os.Getenv("DB_PASS"),
			Name:  os.Getenv("DB_NAME"),
		}
	)

	if db, err = sql.Open(c.Drive, "host="+c.Host+" port="+c.Port+" user="+c.User+" password="+c.Pass+" dbname="+c.Name+" sslmode=disable"); err != nil {
		return t, err
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		return t, err
	}

	transaction, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  readOnly,
	})

	if err != nil {
		return nil, err
	}

	t.ctx = ctx
	t.postgres = transaction
	t.Builder = squirrel.StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		RunWith(t.postgres)

	return t, nil
}

// Commit commit pending transactions for all open databases
func (t *DBTransaction) Commit() (erro error) {
	return t.postgres.Commit()
}

// Rollback rollback transaction pending for all open databases
func (t *DBTransaction) Rollback() {
	_ = t.postgres.Rollback()
}
