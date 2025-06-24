package sqlator

import (
	"github.com/flazhgrowth/fg-tamagochi/pkg/db/sqlator/sqlreader"
	"github.com/flazhgrowth/fg-tamagochi/pkg/db/sqlator/sqlwriter"
)

func (actuator *SQLatorImpl) Reader() sqlreader.SQLReader {
	return actuator.reader
}
func (actuator *SQLatorImpl) Writer() sqlwriter.SQLWriter {
	return actuator.writer
}
