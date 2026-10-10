package storage

import (
	"fmt"

	"github.com/dejitarudemon/ignicula/internal/value"
)

// NotFoundError indicates that no value exists for the requested key.
type NotFoundError struct {
	key value.Key
}

// NewNotFoundError returns a [NotFoundError] for key.
func NewNotFoundError(key value.Key) NotFoundError {
	return NotFoundError{
		key: key,
	}
}

// Key returns the missing key.
func (e NotFoundError) Key() value.Key {
	return e.key
}

// Error implements the error interface.
func (e NotFoundError) Error() string {
	return fmt.Sprintf("value not found with %s key", e.key)
}
