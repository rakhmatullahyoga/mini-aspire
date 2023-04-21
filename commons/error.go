package commons

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrInvalidRecord  = errors.New("invalid record")
)
