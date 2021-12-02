package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/Masterminds/squirrel"
)

// DBTransaction used to aggregate transactions
type DBTransaction struct {
	postgres *sql.Tx
	Builder  squirrel.StatementBuilderType
}

var (
	host     = os.Getenv("DB_HOST")
	port, _  = strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASS")
	dbname   = os.Getenv("DB_NAME")
)

// OpenConnection initialize connection with database
func OpenConnection(readOnly bool) (*DBTransaction, error) {
	t := &DBTransaction{}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	ctx, err := db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  readOnly,
	})

	if err != nil {
		return nil, err
	}

	t.postgres = ctx
	t.Builder = squirrel.StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		RunWith(t.postgres)

	return t, nil
}

// Commit commit pending transactions for all open databases
func (t *DBTransaction) Commit() (erro error) {
	erro = t.postgres.Commit()
	return
}

// Rollback rollback transaction pending for all open databases
func (t *DBTransaction) Rollback() {
	_ = t.postgres.Rollback()
}
