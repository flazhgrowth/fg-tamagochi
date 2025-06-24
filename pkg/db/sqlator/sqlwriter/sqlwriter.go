package sqlwriter

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/flazhgrowth/fg-tamagotchi/pkg/db/sqlator/sqltx"
	"github.com/jmoiron/sqlx"
)

type SQLWriter interface {
	// Write does destructive operations, where it can be insert, update, or delete data.
	/*
		The method itself is pretty general, so it accepts query and args. This method can be reference as example to standardize the codebase itself.
	*/
	Write(ctx context.Context, query string, args []any) (lastInsertedID int64, err error)

	// Insert does destructive operations on inserting new data to the database
	/*
		Use this method if you need the id created. Last argument of the method is the pointer to the id variable
	*/
	Insert(ctx context.Context, query string, args []any, dest any) (err error)
}

type SQLWriterImpl struct {
	actuator *sqlx.DB
}

func New(db *sqlx.DB) SQLWriter {
	return &SQLWriterImpl{actuator: db}
}

var driverToLastInsertedID map[string]func(r sql.Result) (lastInsertedID int64, err error) = map[string]func(r sql.Result) (lastInsertedID int64, err error){
	"postgres": func(r sql.Result) (lastInsertedID int64, err error) {
		return 0, nil
	},
	"mysql": func(r sql.Result) (lastInsertedID int64, err error) {
		return r.LastInsertId()
	},
}

func (impl *SQLWriterImpl) Write(ctx context.Context, query string, args []any) (lastInsertedID int64, err error) {
	tx, finishFn, isNewTx := sqltx.GetTxFromContext(ctx, impl.actuator)
	if isNewTx {
		defer finishFn(tx, &err)
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	lastInsertedIDHandler, found := driverToLastInsertedID[impl.actuator.DriverName()]
	if !found {
		return 0, nil
	}

	return lastInsertedIDHandler(res)
}

func (impl *SQLWriterImpl) Insert(ctx context.Context, query string, args []any, dest any) (err error) {
	tx, finishFn, isNewTx := sqltx.GetTxFromContext(ctx, impl.actuator)
	if isNewTx {
		defer finishFn(tx, &err)
	}
	query = fmt.Sprintf("%s RETURNING id", query)

	return tx.QueryRowxContext(ctx, query, args...).Scan(dest)
}
