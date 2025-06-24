package sqltx

import (
	"context"

	"github.com/flazhgrowth/fg-tamagochi/constant"
	"github.com/jmoiron/sqlx"
)

type FnTxFinisher func(tx *sqlx.Tx, err *error)

func GetTxFromContext(ctx context.Context, db *sqlx.DB) (tx *sqlx.Tx, finishFn FnTxFinisher, isNewTx bool) {
	tx, ok := ctx.Value(constant.CtxKeyDBTx).(*sqlx.Tx)
	if !ok {
		return db.MustBeginTx(ctx, nil), finishTx, true
	}

	return tx, finishTx, false
}
