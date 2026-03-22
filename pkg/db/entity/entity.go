package entity

import (
	"database/sql"
	"time"
)

type (
	BaseModel struct {
		ID        uint64       `db:"id"`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
	}
	BaseModelSoftDelete struct {
		BaseModel
		DeletedAt sql.NullTime `db:"deleted_at"`
	}
)

type (
	BaseULIDModel struct {
		ID        string       `db:"id"`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
	}
	BaseULIDModelSoftDelete struct {
		BaseULIDModel
		DeletedAt sql.NullTime `db:"deleted_at"`
	}
)
