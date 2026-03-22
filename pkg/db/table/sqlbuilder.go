package table

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/flazhgrowth/fg-tamagochi/pkg/db/driver"
	"github.com/flazhgrowth/fg-tamagochi/pkg/vault"
)

var (
	errTableHasNotInstantiated error = fmt.Errorf("table has not yet been instantiated")
)

func (table *Table) SelectQuery() (squirrel.SelectBuilder, error) {
	if err := table.checks(); err != nil {
		return squirrel.SelectBuilder{}, err
	}

	builder := squirrel.
		Select(table.SelectColumns...).
		From(table.Name)

	if vault.GetVault().Database.Driver.Is(driver.DriverPostgres) {
		return builder.PlaceholderFormat(squirrel.Dollar), nil
	}

	return builder.PlaceholderFormat(squirrel.Question), nil
}

func (table *Table) CountQuery() (squirrel.SelectBuilder, error) {
	if err := table.checks(); err != nil {
		return squirrel.SelectBuilder{}, err
	}

	builder := squirrel.
		Select(table.CountColumns...).
		From(table.Name)

	if vault.GetVault().Database.Driver.Is(driver.DriverPostgres) {
		return builder.PlaceholderFormat(squirrel.Dollar), nil
	}

	return builder.PlaceholderFormat(squirrel.Question), nil
}

func (table *Table) InsertQuery() (squirrel.InsertBuilder, error) {
	if err := table.checks(); err != nil {
		return squirrel.InsertBuilder{}, err
	}

	builder := squirrel.
		Insert(table.Name)

	if vault.GetVault().Database.Driver.Is(driver.DriverPostgres) {
		return builder.PlaceholderFormat(squirrel.Dollar), nil
	}

	return builder.PlaceholderFormat(squirrel.Question), nil
}

func (table *Table) UpdateQuery() (squirrel.UpdateBuilder, error) {
	if err := table.checks(); err != nil {
		return squirrel.UpdateBuilder{}, err
	}

	builder := squirrel.
		Update(table.Name)

	if vault.GetVault().Database.Driver.Is(driver.DriverPostgres) {
		return builder.PlaceholderFormat(squirrel.Dollar), nil
	}

	return builder.PlaceholderFormat(squirrel.Question), nil
}

func (table *Table) DeleteQuery() (squirrel.DeleteBuilder, error) {
	if err := table.checks(); err != nil {
		return squirrel.DeleteBuilder{}, err
	}

	builder := squirrel.
		Delete(table.Name)

	if vault.GetVault().Database.Driver.Is(driver.DriverPostgres) {
		return builder.PlaceholderFormat(squirrel.Dollar), nil
	}

	return builder.PlaceholderFormat(squirrel.Question), nil
}

func (table *Table) checks() error {
	if table.Name == "" || len(table.SelectColumns) == 0 || len(table.CountColumns) == 0 || len(table.InsertColumns) == 0 {
		return errTableHasNotInstantiated
	}

	return nil
}
