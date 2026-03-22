package sqlwriter

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/flazhgrowth/fg-tamagochi/pkg/db/driver"
	"github.com/flazhgrowth/fg-tamagochi/pkg/db/sqlator/sqltx"
	"github.com/jmoiron/sqlx"
)

type SQLWriter interface {
	// Write does destructive operations, where it can be insert, update, or delete data.
	/*
		The method itself is pretty general, so it accepts query and args. This method can be reference as example to standardize the codebase itself.
		Notes:
			Write by default checks from context if there are any db transaction exists on the context.
			If db transaction exists, it will use that db transaction and refrain to commit/rollback inside the method.
			If, no db transaction found, it will generate new db transaction and commit/rollback once the method completed.
	*/
	Write(ctx context.Context, query string, args []any) (lastInsertedID int64, err error)

	// Insert does destructive operations on inserting new data to the database
	/*
		Use this method if you need the id created. Last argument of the method is the pointer to the id variable / id field of a struct
		Notes:
			Insert also, by default checks from context if there are any db transaction exists on the context.
			If db transaction exists, it will use that db transaction and refrain to commit/rollback inside the method.
			If, no db transaction found, it will generate new db transaction and commit/rollback once the method completed.
	*/
	Insert(ctx context.Context, query string, args []any, dest any) (err error)
}

type writer struct {
	actuator *sqlx.DB
}

func New(db *sqlx.DB) SQLWriter {
	return &writer{actuator: db}
}

type lastInsertFn func(r sql.Result) (lastInsertedID int64, err error)

var driverToLastInsertedID map[driver.Driver]lastInsertFn = map[driver.Driver]lastInsertFn{
	driver.DriverPostgres: func(r sql.Result) (lastInsertedID int64, err error) {
		return 0, nil
	},
	driver.DriverMySql: func(r sql.Result) (lastInsertedID int64, err error) {
		return r.LastInsertId()
	},
}

func (impl *writer) Write(ctx context.Context, query string, args []any) (lastInsertedID int64, err error) {
	tx, finishFn, isNewTx := sqltx.GetTxFromContext(ctx, impl.actuator)
	if isNewTx {
		defer finishFn(tx, &err)
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	lastInsertedIDHandler, found := driverToLastInsertedID[driver.Driver(impl.actuator.DriverName())]
	if !found {
		return 0, nil
	}

	return lastInsertedIDHandler(res)
}

func (impl *writer) Insert(ctx context.Context, query string, args []any, dest any) (err error) {
	tx, finishFn, isNewTx := sqltx.GetTxFromContext(ctx, impl.actuator)
	if isNewTx {
		defer finishFn(tx, &err)
	}
	query = fmt.Sprintf("%s RETURNING id", query)

	return tx.QueryRowxContext(ctx, query, args...).Scan(dest)
}
