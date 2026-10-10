package storage

import (
	"context"

	"github.com/dejitarudemon/ignicula/internal/value"
)

// Storage is a context-aware key-value store for [value.Value] entries.
//
// Get must return a [NotFoundError] when key is absent.
// Put overwrites any existing entry for key.
// Delete is a no-op when key is absent; it must not return [NotFoundError].
type Storage interface {
	Get(ctx context.Context, key value.Key) (value.Value, error)
	Put(ctx context.Context, key value.Key, value value.Value) error
	Delete(ctx context.Context, key value.Key) error
}
