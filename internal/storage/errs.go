package storage

import (
	"fmt"

	"github.com/dejitarudemon/ignicula/internal/value"
)

type NotFoundError struct {
	key value.Key
}

func NewNotFoundError(key value.Key) NotFoundError {
	return NotFoundError{
		key: key,
	}
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("value not found wiht %s key", e.key)
}
