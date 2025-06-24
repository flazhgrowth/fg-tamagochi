package app

import (
	"github.com/flazhgrowth/fg-tamagochi/pkg/db/sqlator"
	"github.com/flazhgrowth/fg-tamagochi/pkg/db/sqlator/sqltx"
)

func (app *App) GetSQLator() sqlator.SQLator {
	return app.sqlator
}

func (app *App) GetTxSQLator() sqltx.SQLTx {
	return app.txsqlator
}
