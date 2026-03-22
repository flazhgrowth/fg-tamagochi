package table

type (
	Table struct {
		Name          string
		SelectColumns []string
		InsertColumns []string
		CountColumns  []string
	}
)
