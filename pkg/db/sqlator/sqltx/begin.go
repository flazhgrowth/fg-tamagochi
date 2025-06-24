package sqltx

import (
	"context"

	"github.com/flazhgrowth/fg-tamagotchi/constant"
)

func (impl *SQLTxImpl) Begin(ctx context.Context) (context.Context, error) {
	tx, err := impl.actuator.BeginTxx(ctx, nil)
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, constant.CtxKeyDBTx, tx), nil
}
