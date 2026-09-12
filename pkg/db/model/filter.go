package model

import (
	"github.com/Masterminds/squirrel"
)

type (
	Filter[T any] struct {
		Valid      bool
		V          T
		Comparator Comparator
	}

	// Use this if you need a direct controlled condition (eg: squirrel.Eq, squirrel.Lt, etc)
	/*
		Notes: it is recommended to use Filter[T any] insteaf of FilterSqlizer, to make the service layer agnostic to the repo layer condition query
	*/
	FilterSqlizer struct {
		Valid     bool
		Condition squirrel.Sqlizer
	}
)

type Comparator string

var (
	Eq    Comparator = "="
	Lt    Comparator = "<"
	Gt    Comparator = ">"
	LtEq  Comparator = "<="
	GtEq  Comparator = ">="
	Like  Comparator = "like"
	ILike Comparator = "ilike"
)

// ConditionQuery apply Where condition to squirrel.SelectBuilder.
/*
	By default, if Comparator is not one of Eq, Lt, Gt, LtEq, GtEq, Like, or ILike, it will directly use squirrel.Eq{}
	This also mean, not passing Comparator will assume it is meant for equal checking
*/
func (filter Filter[T]) ConditionQuery(builder squirrel.SelectBuilder, col string) squirrel.SelectBuilder {
	if !filter.Valid {
		return builder
	}

	switch filter.Comparator {
	case Eq:
		return builder.Where(squirrel.Eq{col: filter.V})
	case Lt:
		return builder.Where(squirrel.Lt{col: filter.V})
	case Gt:
		return builder.Where(squirrel.Gt{col: filter.V})
	case LtEq:
		return builder.Where(squirrel.LtOrEq{col: filter.V})
	case GtEq:
		return builder.Where(squirrel.GtOrEq{col: filter.V})
	case Like:
		return builder.Where(squirrel.Like{col: filter.V})
	case ILike:
		return builder.Where(squirrel.ILike{col: filter.V})
	default:
		return builder.Where(squirrel.Eq{col: filter.V})
	}
}
