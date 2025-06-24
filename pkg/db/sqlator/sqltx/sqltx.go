package sqltx

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var (
	errNilTx error = fmt.Errorf("tx cannot be nil")
)

type SQLTx interface {
	Begin(ctx context.Context) (context.Context, error)
	Finish(ctx context.Context, err *error)
}

type SQLTxImpl struct {
	actuator *sqlx.DB
}

func New(db *sqlx.DB) SQLTx {
	return &SQLTxImpl{
		actuator: db,
	}
}
