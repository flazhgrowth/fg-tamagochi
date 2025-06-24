package sqltx

import (
	"context"

	"github.com/flazhgrowth/fg-tamagotchi/constant"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (impl *SQLTxImpl) Finish(ctx context.Context, err *error) {
	tx := ctx.Value(constant.CtxKeyDBTx).(*sqlx.Tx)
	finishTx(tx, err)
}

func finishTx(tx *sqlx.Tx, err *error) {
	if tx == nil {
		log.Error().Msg(errNilTx.Error())
		return
	}

	var errOg error
	if err != nil {
		errOg = *err
	}

	if errOg != nil {
		if errRb := tx.Rollback(); errRb != nil {
			log.Error().Msgf("error during rollback: %s", errRb.Error())
		}
		log.Error().Msgf("error on transactions: %s", errOg.Error())
		return
	}

	if errCmt := tx.Commit(); errCmt != nil {
		log.Error().Msgf("error during commit: %s", errCmt.Error())
	}
}
