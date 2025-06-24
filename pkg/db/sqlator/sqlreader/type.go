package sqlreader

import "errors"

var (
	errDestNotPointer        error = errors.New("destination must be a pointer")
	errDestNotPointerToSlice error = errors.New("destination must be a pointer to slice")
)
