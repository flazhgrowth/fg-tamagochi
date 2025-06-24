package sqlreader

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/jmoiron/sqlx"
)

type SQLReader interface {
	// Get gets data. Returned data expected to be one data only. Dest must be pointer to val/struct
	Get(ctx context.Context, query string, args []any, dest any) (err error)

	// Find finds data. Returned data can be more than one data. Dest must be pointer to slice
	Find(ctx context.Context, query string, args []any, dest any) (err error)
}

type SQLReaderImpl struct {
	actuator *sqlx.DB
}

func New(db *sqlx.DB) SQLReader {
	return &SQLReaderImpl{actuator: db}
}

func (impl *SQLReaderImpl) Get(ctx context.Context, query string, args []any, dest any) (err error) {
	row := impl.actuator.QueryRowxContext(ctx, query, args...)
	destVal := reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Pointer {
		return errDestNotPointer
	}
	elemVal := destVal.Elem()
	if elemVal.Kind() == reflect.Struct {
		if err = row.StructScan(dest); err != nil {
			return err
		}
	} else {
		if err = row.Scan(dest); err != nil {
			return err
		}
	}

	return nil
}

func (impl *SQLReaderImpl) Find(ctx context.Context, query string, args []any, dest any) (err error) {
	rows, err := impl.actuator.QueryxContext(ctx, query, args...)
	if err != nil {
		return err
	}
	// dest must be a pointer to a slice
	destVal := reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Pointer {
		return errDestNotPointer
	}

	sliceVal := destVal.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return errDestNotPointerToSlice
	}

	elemType := sliceVal.Type().Elem()
	for rows.Next() {
		elemPtr := reflect.New(elemType)
		if err := rows.StructScan(elemPtr.Interface()); err != nil {
			return err
		}
		sliceVal.Set(reflect.Append(sliceVal, elemPtr.Elem()))
	}

	if sliceVal.Len() == 0 {
		return sql.ErrNoRows
	}

	return nil
}
