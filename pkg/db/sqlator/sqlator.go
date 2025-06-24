package sqlator

import (
	"github.com/flazhgrowth/fg-tamagotchi/pkg/db/sqlator/sqlreader"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/db/sqlator/sqltx"
	"github.com/flazhgrowth/fg-tamagotchi/pkg/db/sqlator/sqlwriter"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SQLator interface {
	Writer() sqlwriter.SQLWriter
	Reader() sqlreader.SQLReader
}

type SQLatorImpl struct {
	reader sqlreader.SQLReader
	writer sqlwriter.SQLWriter
}

func New(cfg SQLatorConfig) (SQLator, sqltx.SQLTx) {
	writerDB := sqlx.MustOpen(cfg.Driver, cfg.WriterDSN)
	readerDB := sqlx.MustOpen(cfg.Driver, cfg.ReaderDSN)

	return &SQLatorImpl{
		writer: sqlwriter.New(writerDB),
		reader: sqlreader.New(readerDB),
	}, sqltx.New(writerDB)
}
